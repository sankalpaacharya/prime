package main

import (
	"fmt"
	"prime/lexer"
)

func main() {
	input := "+;,)(}{"
	cal := lexer.New(input)

	fmt.Println(cal.NextToken())
	fmt.Println(cal.NextToken())

}
