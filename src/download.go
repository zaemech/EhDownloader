package main

import (
    "fmt"
    "io"
    "bufio"
    "os"
    "regexp"
    "strings"
    "strconv"
    "math"
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


func count_pages(rootPage string) int {
    regex := regexp.MustCompile(`Showing \d+ - \d+ of \d+ images`)
    result := regex.FindString(rootPage)

    tmp := strings.Split(result, " ")
    images, _ := strconv.Atoi(tmp[3])
    total,  _ := strconv.Atoi(tmp[5])

    if images == total {
        return 1
    } else if images <= total {
        return int(math.Ceil(float64(total) / float64(images)))
    }

    // more images on the page than there are in the gallery?
    return 0
}


func clean_url(dirtyUrl string) string {
    regex := regexp.MustCompile(`\d{6}/[\da-f]{10}`)
    tmp := regex.FindString(dirtyUrl)

    if tmp == "" {
        return ""
    }

    // eventually work with exhentai too
    url := "http://g.e-hentai.org/g/" + tmp + "/"
    return url
}


func download(args []string) {
    if len(args) < 1 {
        usage()
        return
    }

    rootPage := load_url(clean_url(args[0]))
    numPage := count_pages(rootPage)
    fmt.Println(numPage)
}
