package main

import (
    "fmt"
    "io"
    "bufio"
    "os"
)


func rosebud(err error) {
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }
}


func read_in_queue() []string {
    var links []string

    file, err := os.Open("queue.txt")
    rosebud(err)

    reader := bufio.NewReader(file)

    for {
        line, err := reader.ReadString('\n')
        links = append(links, line)

        if err == io.EOF {
            break
        } else {
            rosebud(err)
        }
    }

    file.Close()
    return links
}


func main() {
    links := read_in_queue()

    for i := range links {
        fmt.Printf("%s", links[i])
    }
}
