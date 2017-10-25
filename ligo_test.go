package ligo

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeySearch(t *testing.T) {
	byteData := "{ \"number\": 123, \"object\": { \"child\": \"hello world!\" } }"
	var data interface{}
	err := json.Unmarshal([]byte(byteData), &data)
	assert.Nil(t, err)
	integer := KeySearch("number", data, false).(float64)
	assert.Equal(t, integer, float64(123))

	child := KeySearch("child", data, false).(string)
	assert.Equal(t, child, "hello world!")

	obj := KeySearch("object", data, true).(map[string]interface{})
	assert.Equal(t, obj, map[string]interface{}{"child": "hello world!"})

}

func TestCheckType(t *testing.T) {
	i := interface{}(123)
	y := CheckType(i)
	assert.Equal(t, y, "int")
	i = interface{}("string")
	y = CheckType(i)
	assert.Equal(t, y, "string")
}
