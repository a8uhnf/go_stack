package test

import "fmt"

type Book struct {
	Title string
	Author string
	Pages int64
}

func (ko Book) CategoryByLength() string {
	var ret string
	if ko.Pages > 300 {
		ret = "NOVEL"
	} else {
		ret = "SHORT STORY"
	}
	return ret
}

func hello(s string) (string) {
	return fmt.Sprintf("Hello, %s", s)
}
