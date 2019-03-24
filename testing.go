package main

import (
	"fmt"
	"url-shorter/util"
)

func main() {

	isvalidate := util.Shortener{}.ValidateUrl("abc")

	fmt.Println(isvalidate)

	randomString := util.Shortener{}.GetRandomString()

	fmt.Println(randomString)
}
