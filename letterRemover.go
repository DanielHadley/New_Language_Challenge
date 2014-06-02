package main

import (
        "fmt"
        "log"
        "regexp"
)

func main() {
    //The string followed by the + is what you will replace
        reg, err := regexp.Compile("(?i)"+"A") 
        if err != nil {
                log.Fatal(err)
        }

        safe := reg.ReplaceAllString("Snarf", "") 
        fmt.Println(safe)   // Output: a*-+fe5v9034,j*.ae6
}