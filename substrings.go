/*
Outputs all the possible substrings(subsets) of a string(slice of bytes) where order does not matter.
*/
package main

import (
	"log"
)

//Takes in a string, a temp string accumulator, and results set where the final
//result is stored.  It runs in O(2^n), and is generally expensive.
func SubSets(s string, acc string, results *[]string) {

	if len(s) == 0 {
		*results = append(*results, acc)
		return
	}

	SubSets(string(s[1:len(s)]), acc+string(s[0]), sets)
	SubSets(string(s[1:len(s)]), acc, sets)

}

func main() {

	results := make([]string, 0)
	SubSets("ask", "", &results)
	log.Println(results)

}
