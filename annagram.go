package main

import (
    "fmt"
    "strings"
    "sort"
)

func annagramFinder(a, b string) bool {
    c := strings.Split(a, "") 
    sort.Strings(c)
    d := strings.Split(b, "") 
    sort.Strings(d)
    if len(c) != len(d) {
        return false
    }
    for i, v := range c {
        if v != d[i] {
            return false
        }
    }
    return true
}


func main() {
	p := (annagramFinder("test", "sett"))
    fmt.Printf("%q\n", p)
}
