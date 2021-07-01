package bencode

import (
	"bytes"
	"fmt"
)

// encoder ...
type encoder struct {
	w *bytes.Buffer
}

// newEncoder ...
func newEncoder() *encoder {
	return &encoder{
		w: &bytes.Buffer{},
	}
}

// encodeList ...
func (e *encoder) encodeList(l []interface{}) error {
	if err := e.w.WriteByte('l'); err != nil {
		return err
	}
	for _, v := range l {
		if _, err := e.encode(v); err != nil {
			return err
		}
	}
	if err := e.w.WriteByte('e'); err != nil {
		return err
	}
	return nil
}

// encodeDict ...
func (e *encoder) encodeDict(m map[string]interface{}) error {
	if err := e.w.WriteByte('d'); err != nil {
		return err
	}
	for k, v := range m {
		_, err := fmt.Fprintf(e.w, "%d:%s", len(k), k)
		if err != nil {
			return err
		}
		_, err = e.encode(v)
		if err != nil {
			return err
		}
	}
	if err := e.w.WriteByte('e'); err != nil {
		return err
	}
	return nil
}

// encode ...
func (e *encoder) encode(v interface{}) ([]byte, error) {
	var (
		err error
	)
	switch val := v.(type) {
	case int, int8, int16, int32, int64:
		_, err = fmt.Fprintf(e.w, "i%de", val)
	case uint, uint8, uint16, uint32, uint64:
		_, err = fmt.Fprintf(e.w, "i%ue", val)
	case bool:
		i := 0
		if val {
			i = 1
		}
		_, err = fmt.Fprintf(e.w, "i%de", i)
	case string:
		_, err = fmt.Fprintf(e.w, "%d:%s", len(val), val)
	case []byte:
		_, err = fmt.Fprintf(e.w, "%d:", len(val))
		if err != nil {
			return nil, err
		}
		_, err = e.w.Write(val)
	case []interface{}:
		err = e.encodeList(val)
	case map[string]interface{}:
		err = e.encodeDict(val)
	}
	if err != nil {
		return nil, err
	}
	return e.w.Bytes(), nil
}

// Encode ...
func Encode(v interface{}) ([]byte, error) {
	return newEncoder().encode(v)
}
