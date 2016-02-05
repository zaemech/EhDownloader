package main

import (
    "fmt"
//    "io"
//    "io/ioutil"
//    "bufio"
    "os"
//    "regexp"
//    "net/http"
//    "strings"
//    "strconv"
)


func rosebud(err error) {
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }
}


func main() {
	if len(os.Args) < 2 {
        fmt.Printf("Usage information to come.\n")
        // display usage
        return
    }

    args := os.Args[1:]

    if args[0] == "download" {
        fmt.Printf("download not implemented.\n")
        //download(args[1:])
    } else if args[0] == "info" {
        info(args[1:])
    } else {
        fmt.Printf("Usage information to come.\n")
        // display usage
    }
}
