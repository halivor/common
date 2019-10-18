package util

import "testing"

func TestToken() {
	key := "P6IyJq6hQD8jwZ9I"
	text := "P6IyJq6hQD8jwZ9I1234567890123456"
	ed := u.EncData([]byte(key), []byte(text))
	fmt.Println(ed)
	bed, _ := hex.DecodeString(ed)
	dd := u.DecData([]byte(key), bed)
	fmt.Println(dd)
}
