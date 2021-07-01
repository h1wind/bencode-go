package bencode_test

import (
	"fmt"
	"testing"

	"github.com/h1zzz/bencode-go"
)

func TestDecode(t *testing.T) {
	data := []byte("d4:porti6881e1:ad6:target20:mnopqrstuvwxyz1234562:id20:abcdefghij0123456789e1:y1:e1:t2:aa1:eli201e23:A Generic Error Ocurredee")
	r, err := bencode.Decode(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", r)
}

func TestEncode(t *testing.T) {
	msg := map[string]interface{}{
		"a": map[string]interface{}{
			"id":     "abcdefghij0123456789",
			"target": "mnopqrstuvwxyz123456",
		},
		"e": []interface{}{
			201,
			"A Generic Error Ocurred",
		},
		"port": 6881,
		"t":    "aa",
		"y":    "e",
	}
	data, err := bencode.Encode(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", string(data))
	r, err := bencode.Decode(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", r)
}
