package main

import (
	"fmt"
	"strings"
)


type Link struct {
	Url      string
	Depth    int
	Children []*Link
}

func (l Link) toString() string {
	var builder strings.Builder
	printLink(&builder, l, 0)
	return builder.String()
}

func printLink(builder *strings.Builder, link Link, indent int) {
	// Add indentation based on depth
	for i := 0; i < indent; i++ {
		builder.WriteString("|   ")
	}

	// Append link information
	builder.WriteString(fmt.Sprintf("%s\n", link.Url))

	// Recursively print children
	for _, child := range link.Children {
		printLink(builder, *child, indent+1)
	}
}
