package main

import (
    "fmt"
    "os"
)


func rosebud(err error) {
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }
}

func usage() {
    fmt.Println("Download gallery\n",
                "  $ download [url]\n",
                "Download a list of galleries (NOT IMPLEMENTED)\n",
                "  $ download -f [file]\n")
}

func main() {
    if len(os.Args) < 2 {
        usage()
        return
    }

    args := os.Args[1:]

    if args[0] == "download" {
        download(args[1:])
    } else if args[0] == "info" {
        //info(args[1:])
    } else {
        usage()
    }
}
