/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTechnicalAssetTypeMarshalJSON(t *testing.T) {
	// Test individual enum marshaling
	for _, assetType := range []TechnicalAssetType{ExternalEntity, Process, Datastore} {
		data, err := json.Marshal(assetType)
		assert.NoError(t, err)
		assert.Equal(t, "\""+assetType.String()+"\"", string(data))
	}

	// Test marshaling as part of a struct
	type TestStruct struct {
		Type TechnicalAssetType `json:"type"`
	}

	testCases := map[string]struct {
		input    TestStruct
		expected string
	}{
		"external-entity": {
			input:    TestStruct{Type: ExternalEntity},
			expected: `{"type":"external-entity"}`,
		},
		"process": {
			input:    TestStruct{Type: Process},
			expected: `{"type":"process"}`,
		},
		"datastore": {
			input:    TestStruct{Type: Datastore},
			expected: `{"type":"datastore"}`,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			data, err := json.Marshal(testCase.input)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expected, string(data))
		})
	}
}

func TestTechnicalAssetTypeUnmarshalJSON(t *testing.T) {
	// Test unmarshaling as part of a struct
	type TestStruct struct {
		Type TechnicalAssetType `json:"type"`
	}

	testCases := map[string]struct {
		input    string
		expected TestStruct
	}{
		"external-entity": {
			input:    `{"type":"external-entity"}`,
			expected: TestStruct{Type: ExternalEntity},
		},
		"process": {
			input:    `{"type":"process"}`,
			expected: TestStruct{Type: Process},
		},
		"datastore": {
			input:    `{"type":"datastore"}`,
			expected: TestStruct{Type: Datastore},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			var actual TestStruct
			err := json.Unmarshal([]byte(testCase.input), &actual)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
