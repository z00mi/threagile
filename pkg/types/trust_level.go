package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type TrustLevel int

const (
	UndefinedTrustLevel TrustLevel = iota
	ValidUserCredentials
	ApiKey
	ServiceAccount
)

func TrustLevelValues() []TypeEnum {
	return []TypeEnum{
		UndefinedTrustLevel,
		ValidUserCredentials,
		ApiKey,
		ServiceAccount,
	}
}

func ParseTrustLevel(value string) (trustLevel TrustLevel, err error) {
	if len(value) == 0 {
		return UndefinedTrustLevel, nil
	}
	return TrustLevel(0).Find(value)
}

var TrustLevelTypeDescription = [...]TypeDescription{
	{"", "Undefined"}, // UndefinedDataClassification
	{"valid-user-credentials", "Valid User Credentials"},
	{"api-key", "API Key"},
	{"service-account", "Service Account"},
}

func (what TrustLevel) String() string {
	// NOTE: maintain list also in schema.json for validation in IDEs
	return TrustLevelTypeDescription[what].Name
}

func (what TrustLevel) Explain() string {
	return TrustLevelTypeDescription[what].Description
}

// func (what TrustLevel) AttackerAttractivenessForAsset() float64 {
// 	// fibonacci starting at 1
// 	return [...]float64{1, 2, 3, 5, 8}[what]
// }
// func (what TrustLevel) AttackerAttractivenessForProcessedOrStoredData() float64 {
// 	// fibonacci starting at 1
// 	return [...]float64{1, 2, 3, 5, 8}[what]
// }
// func (what TrustLevel) AttackerAttractivenessForInOutTransferredData() float64 {
// 	// fibonacci starting at 1
// 	return [...]float64{1, 2, 3, 5, 8}[what]
// }

// func (what TrustLevel) RatingStringInScale() string {
// 	result := "(rated "
// 	if what == NoDataStored {
// 		result += "1"
// 	}
// 	if what == PublicData {
// 		result += "2"
// 	}
// 	if what == InternalUseData {
// 		result += "3"
// 	}
// 	if what == ConfidentialData {
// 		result += "4"
// 	}
// 	if what == RetrictedData {
// 		result += "5"
// 	}

// 	result += " in scale of 5)"
// 	return result
// }

func (what TrustLevel) Find(value string) (TrustLevel, error) {
	for index, description := range TrustLevelTypeDescription {
		if strings.EqualFold(value, description.Name) {
			return TrustLevel(index), nil
		}
	}

	return TrustLevel(0), fmt.Errorf("unknown DataClassification value %q", value)
}

func (what TrustLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(what.String())
}

func (what *TrustLevel) UnmarshalJSON(data []byte) error {
	var text string
	unmarshalError := json.Unmarshal(data, &text)
	if unmarshalError != nil {
		return unmarshalError
	}

	value, findError := what.Find(text)
	if findError != nil {
		return findError
	}

	*what = value
	return nil
}

func (what TrustLevel) MarshalYAML() (interface{}, error) {
	return what.String(), nil
}

func (what *TrustLevel) UnmarshalYAML(node *yaml.Node) error {
	value, findError := what.Find(node.Value)
	if findError != nil {
		return findError
	}

	*what = value
	return nil
}
