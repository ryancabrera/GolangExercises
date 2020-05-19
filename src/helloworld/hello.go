package helloworld

import (
	"rcabrera/exercises/src/morestrings"
)

func Helloworld() string {
	var reversedString string
	reversedString = morestrings.ReverseRunes("!oG ,olleH")
	return reversedString
}
