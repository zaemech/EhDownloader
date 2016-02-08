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
    regex := regexp.MustCompile(`\d{1,7}/[\da-f]{10}`)
    tmp := regex.FindString(dirtyUrl)

    if tmp == "" {
        return ""
    }

    // eventually work with exhentai too
    url := "http://g.e-hentai.org/g/" + tmp + "/"
    return url
}


func get_img_pages(rootUrl string) []string {
    var imgPages []string

    rootPage := load_url(rootUrl)
    numPages := count_pages(rootPage)

    imgPageRegex := regexp.MustCompile(
        `http://g.e-hentai.org/s/[\da-f]{10}/\d{1,7}-\d+`)

    // TODO: ignore pages referenced in the comments.
    // Probably something like getting the number of images on
    // the page and replacing the -1 in the next line with that.
    imgPages = imgPageRegex.FindAllString(rootPage, -1)

    for page := 1; page <= numPages - 1; page++ {
        tmpUrl := rootUrl + "?p=" + strconv.Itoa(page)
        rootPage = load_url(tmpUrl)
        tmpPages := imgPageRegex.FindAllString(rootPage, -1)
        imgPages = append(imgPages, tmpPages...)
    }

    return imgPages
}


func get_img_urls(imgPages []string) []string {
    var imgUrls []string
    imgRegex := regexp.MustCompile(
        `http://(\d{1,3}[\.\:\/]){4}[0-9]{0,5}.*?\.(jpg||png||gif)`)

    for _, url := range imgPages {
        page := load_url(url)
        tmpUrl := imgRegex.FindString(page)
        imgUrls = append(imgUrls, tmpUrl)
    }
    return imgUrls
}


func download(args []string) {
    if len(args) < 1 {
        usage()
        return
    }

    rootUrl := clean_url(args[0])
    if rootUrl == "" {
        fmt.Printf("Malformed gallery url")
        os.Exit(1)
    }

    imgPages := get_img_pages(rootUrl)
    imgUrls  := get_img_urls(imgPages)
}
