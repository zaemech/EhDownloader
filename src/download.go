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


def download() {
    links := read_in_queue()
    imgre := regexp.MustCompile(
        `http://g.e-hentai.org/s/[\da-f]{10}/\d{6}-\d+`)

    for i := range links {
        if len(links[i]) == 0 {
            break
        }

        rootPage := load_url(links[i])
        numPage := determine_num_pages(rootPage)

        for p := 0; p <= numPage; p++ {
            imageUrls := imgre.FindAllString(string(rootPage), -1)

            for q := range imageUrls {
                download_image(imageUrls[q])
            }

            if p < numPage {
                galleryPage := strings.TrimSpace(links[i])
                galleryPage = strings.Join([]string{galleryPage, "?p="}, "")
                galleryPage += strconv.Itoa(p+1)
                rootPage = load_url(galleryPage)
            }
        }
    }
}
