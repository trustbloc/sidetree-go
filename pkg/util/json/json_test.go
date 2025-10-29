/*
Copyright Gen Digital Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package json

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testData struct {
	FieldC string
	FieldB int
	FieldA string
}

func TestMarshalCanonical(t *testing.T) {
	value1 := &testData{
		FieldC: "valueC_1",
		FieldB: 100,
		FieldA: "valueA_1",
	}
	value2 := &testData{
		FieldC: "valueC_2",
		FieldB: 200,
		FieldA: "valueA_2",
	}

	t.Run("Struct", func(t *testing.T) {
		v1, err := MarshalCanonical(value1)
		require.NoError(t, err)
		assert.NotNil(t, v1)
		fmt.Printf("%s\n", v1)

		v := &testData{}
		err = json.Unmarshal(v1, v)
		require.NoError(t, err)

		require.Equal(t, value1, v)
	})

	t.Run("Array", func(t *testing.T) {
		arr := []*testData{value1, value2}
		v1, err := MarshalCanonical(arr)
		require.NoError(t, err)
		assert.NotNil(t, v1)
		fmt.Printf("%s\n", v1)

		var v []*testData
		err = json.Unmarshal(v1, &v)
		require.NoError(t, err)

		require.Equal(t, arr, v)
	})

	t.Run("Marshal struct error", func(t *testing.T) {
		originalMarshalJSONMap := marshalJSONMap
		defer func() { marshalJSONMap = originalMarshalJSONMap }()

		marshalJSONMap = func(m map[string]interface{}) ([]byte, error) {
			return nil, fmt.Errorf("injected marshal error")
		}

		_, err := MarshalCanonical(value1)
		require.Error(t, err)
	})

	t.Run("Unmarshal struct error", func(t *testing.T) {
		originalUnmarshalJSONMap := unmarshalJSONMap
		defer func() { unmarshalJSONMap = originalUnmarshalJSONMap }()

		unmarshalJSONMap = func(bytes []byte) (map[string]interface{}, error) {
			return nil, fmt.Errorf("injected unmarshal error")
		}

		_, err := MarshalCanonical(value1)
		require.Error(t, err)
	})

	t.Run("Marshal array error", func(t *testing.T) {
		originalMarshalJSONArray := marshalJSONArray
		defer func() { marshalJSONArray = originalMarshalJSONArray }()

		marshalJSONArray = func(a []map[string]interface{}) ([]byte, error) {
			return nil, fmt.Errorf("injected marshal error")
		}

		_, err := MarshalCanonical([]*testData{value1, value2})
		require.Error(t, err)
	})

	t.Run("Unmarshal array error", func(t *testing.T) {
		originalUnmarshalJSONArray := unmarshalJSONArray
		defer func() { unmarshalJSONArray = originalUnmarshalJSONArray }()

		unmarshalJSONArray = func(bytes []byte) ([]map[string]interface{}, error) {
			return nil, fmt.Errorf("injected unmarshal error")
		}

		_, err := MarshalCanonical([]*testData{value1, value2})
		require.Error(t, err)
	})
}

func TestMarshalIndentCanonical(t *testing.T) {
	value1 := &testData{
		FieldC: "valueC_1",
		FieldB: 100,
		FieldA: "valueA_1",
	}

	t.Run("Success", func(t *testing.T) {
		v1, err := MarshalIndentCanonical(value1, "", " ")
		require.NoError(t, err)
		assert.NotNil(t, v1)
		fmt.Printf("%s\n", v1)
	})

	t.Run("Marshal error", func(t *testing.T) {
		originalMarshalJSONMap := marshalJSONMap
		defer func() { marshalJSONMap = originalMarshalJSONMap }()

		marshalJSONMap = func(m map[string]interface{}) ([]byte, error) {
			return nil, fmt.Errorf("injected marshal error")
		}

		_, err := MarshalIndentCanonical(value1, "", " ")
		require.Error(t, err)
	})
}

func TestGetCanonicalContent(t *testing.T) {
	t.Run("Struct", func(t *testing.T) {
		value1 := []byte(`{"field1":"value1","field2":"value2"}`)
		value2 := []byte(`{"field2":"value2","field1":"value1"}`)

		v1, err := getCanonicalContent(value1)
		require.NoError(t, err)
		assert.NotNil(t, v1)

		v2, err := getCanonicalContent(value2)
		require.NoError(t, err)
		assert.Equal(t, v1, v2)
	})

	t.Run("Array", func(t *testing.T) {
		value1 := []byte(`[{"field1":"value1_1","field2":"value2_1"},{"field1":"value1_2","field2":"value2_2"}]`)
		value2 := []byte(`[{"field2":"value2_1","field1":"value1_1"},{"field2":"value2_2","field1":"value1_2"}]`)

		v1, err := getCanonicalContent(value1)
		require.NoError(t, err)
		assert.NotNil(t, v1)

		v2, err := getCanonicalContent(value2)
		require.NoError(t, err)
		assert.Equal(t, v1, v2)
	})

	t.Run("Marshal error", func(t *testing.T) {
		originalMarshalJSONMap := marshalJSONMap
		defer func() { marshalJSONMap = originalMarshalJSONMap }()

		marshalJSONMap = func(m map[string]interface{}) ([]byte, error) {
			return nil, fmt.Errorf("marshal error")
		}

		value1 := []byte(`{"field1":"value1","field2":"value2"}`)
		_, err := getCanonicalContent(value1)
		require.Error(t, err)
	})
}
