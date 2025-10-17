package zaprpc

import (
	"encoding/gob"
	"encoding/json"
	"io"
)

// Codec defines the interface for all encoding/decoding implementations.
type Codec interface {
	Name() string
	Marshal(w io.Writer, v any) error
	Unmarshal(r io.Reader, v any) error
}

// --------------------
// GOB Codec
// --------------------

type GOBCodec struct{}

// Name returns the codec name.
func (g *GOBCodec) Name() string {
	return "gob"
}

// Marshal encodes the given value into the writer using GOB.
func (g *GOBCodec) Marshal(w io.Writer, v any) error {
	enc := gob.NewEncoder(w)
	return enc.Encode(v)
}

// Unmarshal decodes from the reader into the given value using GOB.
func (g *GOBCodec) Unmarshal(r io.Reader, v any) error {
	dec := gob.NewDecoder(r)
	return dec.Decode(v)
}

// --------------------
// JSON Codec
// --------------------

type JSONCodec struct{}

// Name returns the codec name.
func (j *JSONCodec) Name() string {
	return "json"
}

// Marshal encodes the given value into the writer using JSON.
func (j *JSONCodec) Marshal(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v)
}

// Unmarshal decodes from the reader into the given value using JSON.
func (j *JSONCodec) Unmarshal(r io.Reader, v any) error {
	dec := json.NewDecoder(r)
	return dec.Decode(v)
}

