package main

import (
	"github.com/golang/protobuf/jsonpb"
	"null_value"
	"testing"
)
func TestNullValue(t *testing.T) {
	jsonPayload := `
{
	"parameter": {
		"key": "$param",
		"null_value": null
	}
}
`
	foo := &null_value.Foo{}
	err := jsonpb.UnmarshalString(jsonPayload, foo)
	if err != nil {
		t.Errorf("Failed to unmarshall job payload %v", err)
	}
	params := foo.GetParameter()
	if params.GetKey() != "$param" {
		t.Errorf("Failed to unmarshall foo/parameter/$param")
	}

	if _, ok := params.ValueByType.(*null_value.KeyValue_NullValue); !ok {
		t.Errorf("Failed to unmarshall foo/parameter/null_value: expected: KeyValue_NullValue; got: %v", params.GetValueByType())
	}
}
