package issuer

import (
	"encoding/base64"
	"fmt"
	"reflect"

	"github.com/hyperledger/aries-framework-go/component/models/sdjwt/common"
)

type SDJWTBuilderV5 struct {
	debugMode bool
}

func (s *SDJWTBuilderV5) keysToExclude() []string {
	return []string{
		"iss",
		"iat",
		"nbf",
		"exp",
		"cnf",
		"type",
		"status",
		"sub",
	}
}

func NewSDJWTBuilderV5() *SDJWTBuilderV5 {
	return &SDJWTBuilderV5{}
}

func (s *SDJWTBuilderV5) isAlwaysInclude(curPath string, opts *newOpts) bool {
	if opts == nil || len(opts.alwaysInclude) == 0 {
		return false
	}

	_, ok := opts.alwaysInclude[curPath]
	return ok
}

func (s *SDJWTBuilderV5) isIgnored(curPath string, opts *newOpts) bool {
	if opts == nil || len(opts.nonSDClaimsMap) == 0 {
		return false
	}

	_, ok := opts.nonSDClaimsMap[curPath]
	return ok
}

func (s *SDJWTBuilderV5) isRecursive(curPath string, opts *newOpts) bool {
	if opts == nil || len(opts.recursiveClaimMap) == 0 {
		return false
	}

	_, ok := opts.recursiveClaimMap[curPath]
	return ok
}

func (s *SDJWTBuilderV5) extractValueOptions(curPath string, opts *newOpts) valueOption {
	return valueOption{
		IsStructured:    opts.structuredClaims,
		IsAlwaysInclude: s.isAlwaysInclude(curPath, opts),
		IsIgnored:       s.isIgnored(curPath, opts),
		IsRecursive:     s.isRecursive(curPath, opts),
	}
}

type valueOption struct {
	IsStructured    bool
	IsAlwaysInclude bool
	IsIgnored       bool
	IsRecursive     bool
}

func (s *SDJWTBuilderV5) CreateDisclosuresAndDigests(
	path string,
	claims map[string]interface{},
	opts *newOpts,
) ([]string, map[string]interface{}, error) {
	digestsMap := map[string]interface{}{}
	finalSDDigest, err := createDecoyDisclosures(opts)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create decoy disclosures: %w", err)
	}

	var allDisclosures []string
	for key, value := range claims {
		curPath := key
		if path != "" {
			curPath = path + "." + key
		}

		kind := reflect.TypeOf(value).Kind()

		valOption := s.extractValueOptions(curPath, opts)
		switch kind {
		case reflect.Map:
			if valOption.IsIgnored {
				digestsMap[key] = value
			} else if valOption.IsRecursive {
				nestedDisclosures, nestedDigestsMap, mapErr := s.CreateDisclosuresAndDigests(
					curPath,
					value.(map[string]interface{}),
					opts,
				)
				if mapErr != nil {
					return nil, nil, mapErr
				}

				disclosure, disErr := s.createDisclosure(key, nestedDigestsMap, opts)
				if disErr != nil {
					return nil, nil, fmt.Errorf(
						"create disclosure for recursive disclosure value with path [%v]: %w",
						path, disErr)
				}

				finalSDDigest = append(finalSDDigest, disclosure)
				allDisclosures = append(allDisclosures, nestedDisclosures...)
			} else if valOption.IsAlwaysInclude || valOption.IsStructured {
				nestedDisclosures, nestedDigestsMap, mapErr := s.CreateDisclosuresAndDigests(
					curPath,
					value.(map[string]interface{}),
					opts,
				)
				if mapErr != nil {
					return nil, nil, mapErr
				}

				digestsMap[key] = nestedDigestsMap

				allDisclosures = append(allDisclosures, nestedDisclosures...)
			} else { // plain
				disclosure, disErr := s.createDisclosure(key, value, opts)
				if disErr != nil {
					return nil, nil, fmt.Errorf("create disclosure for map object [%v]: %w",
						path, disErr)
				}

				finalSDDigest = append(finalSDDigest, disclosure)
			}
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			if valOption.IsIgnored { // whole array ignored
				digestsMap[key] = value
				continue
			}

			elementsDigest, elementsDisclosures, arrayElemErr := s.processArrayElements(value, curPath, opts)
			if arrayElemErr != nil {
				return nil, nil, arrayElemErr
			}

			if valOption.IsAlwaysInclude || valOption.IsStructured {
				digestsMap[key] = elementsDigest
			} else { // plain
				disclosure, disErr := s.createDisclosure(key, value, opts)
				if disErr != nil {
					return nil, nil, fmt.Errorf("create disclosure for whole err with path [%v]: %w",
						path, disErr)
				}

				finalSDDigest = append(finalSDDigest, disclosure)
			}

			allDisclosures = append(allDisclosures, elementsDisclosures...)
		default:
			disclosure, disErr := s.createDisclosure(key, value, opts)
			if disErr != nil {
				return nil, nil, fmt.Errorf("create disclosure for simple value with path [%v]: %w",
					path, disErr)
			}

			finalSDDigest = append(finalSDDigest, disclosure)
		}
	}

	digests, err := createDigests(finalSDDigest, opts)
	if err != nil {
		return nil, nil, err
	}

	digestsMap[common.SDKey] = digests

	return append(finalSDDigest, allDisclosures...), digestsMap, nil
}

func (s *SDJWTBuilderV5) processArrayElements(
	value interface{},
	path string,
	opts *newOpts,
) ([]interface{}, []string, error) {
	valSl := reflect.ValueOf(value)
	var digestArr []interface{}
	var elementsDisclosures []string
	for i := 0; i < valSl.Len(); i++ {
		elementPath := fmt.Sprintf("%v[%v]", path, i)
		elementOptions := s.extractValueOptions(elementPath, opts)
		elementValue := valSl.Index(i).Interface()

		if elementOptions.IsIgnored {
			digestArr = append(digestArr, elementValue)
			continue
		}

		digest, err := s.createDisclosure("", elementValue, opts)
		if err != nil {
			return nil, nil,
				fmt.Errorf("create element disclosure for path [%v]: %w", elementPath, err)
		}
		elementsDisclosures = append(elementsDisclosures, digest)
		digestArr = append(digestArr, map[string]string{"...": digest})
	}

	return digestArr, elementsDisclosures, nil
}

func (s *SDJWTBuilderV5) createDisclosure(
	key string,
	value interface{},
	opts *newOpts,
) (*DisclosureEntity, error) {
	salt, err := opts.getSalt()
	if err != nil {
		return nil, fmt.Errorf("generate salt: %w", err)
	}

	finalDis := &DisclosureEntity{
		Salt: salt,
	}
	disclosure := []interface{}{salt}
	if key != "" {
		disclosure = append(disclosure, key)
	}
	disclosure = append(disclosure, value)

	disclosureBytes, err := opts.jsonMarshal(disclosure)
	if err != nil {
		return nil, fmt.Errorf("marshal disclosure: %w", err)
	}

	finalDis.Key = key
	finalDis.Value = value
	finalDis.Result = base64.RawURLEncoding.EncodeToString(disclosureBytes)

	if s.debugMode {
		finalDis.DebugArr = disclosure
		finalDis.DebugStr = string(disclosureBytes)
	}

	return finalDis, nil
}

type DisclosureEntity struct {
	Result   string
	Salt     string
	Key      string
	Value    interface{}
	DebugArr []interface{}
	DebugStr string
}
