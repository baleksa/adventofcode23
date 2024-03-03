package main

import (
	"fmt"
	"math/big"
)

func printFunction(function Function) {
	if len(function) == 0 {
		fmt.Println("{}")
		return
	}
	fmt.Println("{")
	printMapping(function[0])
	fmt.Println("")
	for _, mapping := range function[1:] {
		printMapping(mapping)
		fmt.Println("")
	}
	fmt.Println("}")
}

func printMapping(mapping Mapping) {
	printInterval(mapping.x)
	fmt.Print("->")
	printInterval(mapping.y)
	var change, src big.Int
	change.SetUint64(mapping.dst)
	src.SetUint64(mapping.src)
	change.Sub(&change, &src)
	fmt.Printf(" %11d", change.Int64())
}
