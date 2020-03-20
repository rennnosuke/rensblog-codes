package main

import "strings"

func main() {
	var b []byte
	s := "Hello, string to []byte ."
	r := strings.NewReader(s)
	_, err := r.Read(b)
	if err != nil {
		panic("failed to read string as bytes.")
	}
}
