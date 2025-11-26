package main

import (
	"fmt"
	"urlshort/urls"
)

func main() {
	result := urls.Base("urlshort")
	fmt.Println(result)
}
