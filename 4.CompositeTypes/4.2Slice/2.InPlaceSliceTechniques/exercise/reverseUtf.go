package main

import (
	"unicode/utf8"
)

//ReverseUTF reverse utf encoded string
func ReverseUTF(str []byte) {
	//first revers every single rune in place
	for index, runeSize := 0, 0; index < len(str); index += runeSize {
		_, runeSize = utf8.DecodeRune(str[index:])
		for runeByteFrontIndex, runeByteBackIndex := index, index+runeSize-1; runeByteFrontIndex < runeByteBackIndex; runeByteFrontIndex, runeByteBackIndex = runeByteFrontIndex+1, runeByteBackIndex-1 {
			str[runeByteFrontIndex], str[runeByteBackIndex] = str[runeByteBackIndex], str[runeByteFrontIndex]
		}
	}
	//then reverse all the byte slice
	for front, back := 0, len(str)-1; front < back; front, back = front+1, back-1 {
		str[front], str[back] = str[back], str[front]
	}
}
