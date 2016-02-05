package main

import (
    "fmt"
    "io"
    "bufio"
    "os"
    "regexp"
    "strconv"
)

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


func determine_num_pages(rootPage string) int {
    repage := regexp.MustCompile(`sp\(\d+\)`)
    temp := repage.FindAllString(string(rootPage), -1)
    numPage := 0

    if len(temp) >= 1 {
        t := temp[len(temp)-2][3:]
        numPage, _ = strconv.Atoi(t[:len(t)-1])
    }
    return numPage
}




func download(args []string) {
    if len(args) < 1 {
        usage()
        return
    }

    rootPage := load_url(args[0])
    numPage := determine_num_pages(rootPage)
    fmt.Println(numPage)
}
