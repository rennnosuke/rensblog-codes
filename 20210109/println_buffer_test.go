package main

import (
	"fmt"
	"testing"
)

func TestPrintlnBuffer(t *testing.T) {
	s := "こんにちは、Golang"
	b := s[0]
	fmt.Println(b)                  // 227
	fmt.Printf("%v\n", []byte("こ")) // [227 129 147]
	fmt.Printf("%q\n", "こ")         // "12371"
	fmt.Printf("%+q\n", "こ")        // "\u3053"

	r := []rune(s)
	fmt.Println(r[0]) // 12371

	for _, r := range s {
		fmt.Print(string(r))
	}
	
	apple := "🍎"
	rApple := []rune(apple)
	fmt.Println(string(rApple[0]))

	wata := "邊󠄆"
	rWata := []rune(wata)
	fmt.Printf("%+q\n", rWata)
	fmt.Println(len(rWata))
	fmt.Println(rWata[0])
	fmt.Println(string(rWata[0]))
	fmt.Printf("%+q\n",rWata[0])
}
