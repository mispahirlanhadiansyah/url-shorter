package main

import (
	"fmt"
	"url-shorter/util"
)

func main() {

	isvalidate := util.Shortener{}.ValidateUrl("a1b2c3")

	fmt.Println(isvalidate)

	randomString := util.Shortener{}.GetRandomString()

	fmt.Println(randomString)
}
