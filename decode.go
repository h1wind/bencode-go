package bencode

import (
	"errors"
	"fmt"
	"strconv"
)

// decoder ...
type decoder struct {
	data []byte
	pos  int
}

// newDecoder ...
func newDecoder(data []byte) *decoder {
	return &decoder{
		data: data,
		pos:  0,
	}
}

// readBytes ...
func (d *decoder) readBytes(delim byte) []byte {
	var b []byte
	for d.pos < len(d.data) {
		b = append(b, d.data[d.pos])
		if d.data[d.pos] == delim {
			d.pos++
			break
		}
		d.pos++
	}
	return b
}

// readn ...
func (d *decoder) readn(n int) []byte {
	var b []byte
	for i := 0; i < n && d.pos < len(d.data); i++ {
		b = append(b, d.data[d.pos])
		d.pos++
	}
	return b
}

// decodeInt ...
func (d *decoder) decodeInt() (int, error) {
	d.pos++ // skip 'i'
	s := d.readBytes('e')
	l, err := strconv.ParseInt(string(s[:len(s)-1]), 10, 32)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int(l), nil
}

// decodeString ...
func (d *decoder) decodeString() (string, error) {
	// read string length
	s := d.readBytes(':')
	l, err := strconv.ParseInt(string(s[:len(s)-1]), 10, 32)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	s = d.readn(int(l))
	if len(s) != int(l) {
		return "", errors.New("length of the read string is incorrect")
	}
	return string(s), nil
}

// decodeList ...
func (d *decoder) decodeList() ([]interface{}, error) {
	var (
		l   []interface{}
		v   interface{}
		err error
	)
	d.pos++ // skip 'l'
	for d.pos < len(d.data) && d.data[d.pos] != 'e' {
		switch d.data[d.pos] {
		case 'i':
			v, err = d.decodeInt()
		case 'l':
			v, err = d.decodeList()
		case 'd':
			v, err = d.decodeDict()
		default:
			v, err = d.decodeString()
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		l = append(l, v)
	}
	d.pos++ // skip 'e'
	return l, nil
}

// decodeDict ...
func (d *decoder) decodeDict() (map[string]interface{}, error) {
	var (
		k   string
		v   interface{}
		err error
	)
	dict := make(map[string]interface{})
	d.pos++ // skip 'd'
	for d.pos < len(d.data) && d.data[d.pos] != 'e' {
		k, err = d.decodeString()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		switch d.data[d.pos] {
		case 'i':
			v, err = d.decodeInt()
		case 'l':
			v, err = d.decodeList()
		case 'd':
			v, err = d.decodeDict()
		default:
			v, err = d.decodeString()
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		dict[k] = v
	}
	d.pos++ // skip 'e'
	return dict, nil
}

// decode ...
func (d *decoder) decode() (interface{}, error) {
	switch d.data[d.pos] {
	case 'i':
		return d.decodeInt()
	case 'l':
		return d.decodeList()
	case 'd':
		return d.decodeDict()
	default:
		return d.decodeString()
	}
}

// Decode ...
func Decode(data []byte) (interface{}, error) {
	return newDecoder(data).decode()
}
