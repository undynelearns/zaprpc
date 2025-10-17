package zaprpc

import (
	"bytes"
	"reflect"
	"testing"
)

type testStruct struct {
	Name  string
	Value int
}

func TestJSONCodec(t *testing.T) {
	var buf bytes.Buffer
	c := &JSONCodec{}

	orig := testStruct{Name: "example", Value: 123}

	// Marshal
	if err := c.Marshal(&buf, orig); err != nil {
		t.Fatalf("JSON Marshal failed: %v", err)
	}

	// Unmarshal
	var decoded testStruct
	if err := c.Unmarshal(&buf, &decoded); err != nil {
		t.Fatalf("JSON Unmarshal failed: %v", err)
	}

	// Compare
	if !reflect.DeepEqual(orig, decoded) {
		t.Fatalf("Decoded value mismatch.\nGot:  %+v\nWant: %+v", decoded, orig)
	}

	if c.Name() != "json" {
		t.Fatalf("Codec name mismatch: got %q, want %q", c.Name(), "json")
	}
}

func TestGOBCodec(t *testing.T) {
	var buf bytes.Buffer
	c := &GOBCodec{}

	orig := testStruct{Name: "gobtest", Value: 42}

	if err := c.Marshal(&buf, orig); err != nil {
		t.Fatalf("GOB Marshal failed: %v", err)
	}

	var decoded testStruct
	if err := c.Unmarshal(&buf, &decoded); err != nil {
		t.Fatalf("GOB Unmarshal failed: %v", err)
	}

	if !reflect.DeepEqual(orig, decoded) {
		t.Fatalf("Decoded value mismatch.\nGot:  %+v\nWant: %+v", decoded, orig)
	}

	if c.Name() != "gob" {
		t.Fatalf("Codec name mismatch: got %q, want %q", c.Name(), "gob")
	}
}
