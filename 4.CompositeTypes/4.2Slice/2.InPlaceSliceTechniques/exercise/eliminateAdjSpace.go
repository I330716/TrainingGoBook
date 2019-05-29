package main

import (
	"unicode"
	"unicode/utf8"
)

const space rune = ' '

// EliminateAdjasentSpaces squashes each run of adjacent Unicode spaces
//The function work inplace so the passed slice will be changed
func EliminateAdjasentSpaces(str []byte) []byte {
	var bound int //from 0 to bound (exluded) is the final result
	for index, symbolSize := 0, 0; index < len(str); index += symbolSize {
		//thake the size of the next symbol or sequence of white spaces
		symbolSize, isSpaceSequence := spaceSequenceSize(str[index:])
		//we found white spaces sequence
		if isSpaceSequence {
			//squashthe sequnece into sigle ASCII white space
			str[index] = ' '
			bound++
		} else {
			//there is no sequence but it is a single rune
			//copy the rune in place somere back in the underlying array
			// or in the same space if the was no adjacent spaces at the moment
			for copiedByte := 0; copiedByte < symbolSize; copiedByte++ {
				str[bound] = str[index+copiedByte]
				bound++
			}
		}
	}
	return str[:bound]
}

//this function returns
//1.size of the sigle readed rune and false denoting that there is no white spaces sequence
//2.size of the hole sequence of white spaces (more than one) and true
func spaceSequenceSize(str []byte) (int, bool) {
	var spaceSequenceSize int //the result variable
	var isSpaceSequence bool
	for index, symbolSize := 0, 0; index < len(str); index += symbolSize {
		//extract unicode rune and its size
		symbol, symbolSize := utf8.DecodeRune(str[index:])
		//if the rune is white space add the size to the result
		if unicode.IsSpace(symbol) {
			spaceSequenceSize += symbolSize
			if index > 0 {
				isSpaceSequence = true
			}
		} else {
			//withe space sequence ends here if any
			break
		}
	}
	return spaceSequenceSize, isSpaceSequence
}
