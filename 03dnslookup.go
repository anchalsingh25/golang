package main

import (
    "fmt"
    "net"
)

func main() {
    txtrecords, _ := net.LookupTXT("google.com")
    for _, txt := range txtrecords {
        fmt.Println(txt)
    }
}