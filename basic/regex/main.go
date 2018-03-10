package main

import (
	"regexp"
	"fmt"
)

const text = `
My email is laoji@qq.com

`
func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	//re := regexp.MustCompile(".+@.+\\..+")
	// match := re.FindString(text)
	//match := re.FindAllString(text, -1)
	submatch := re.FindAllStringSubmatch(text, -1)
	for _, m := range submatch{
		fmt.Println(m)
	}
}