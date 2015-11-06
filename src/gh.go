package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "bufio"
    "os"
    "regexp"
    "net/http"
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


func load_url(url string) string {
    response, err := http.Get(url)
    rosebud(err)

    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    rosebud(err)
    return string(contents)
}


func load_url_2(url string) []byte {
    // load a file from the drive for testing purposes instead
    file, err := os.Open("test.html")
    rosebud(err)
    defer file.Close()

    rootPage, err := ioutil.ReadAll(file)
    rosebud(err)
    return rootPage
}


func download_image(url string) {
    imgPage := load_url(url)
    re := regexp.MustCompile(
        `http://(\d{1,3}[\.\:\/]){4}[0-9]{0,5}.*?\.(jpg||png||gif)`)

    imgUrl := re.FindString(string(imgPage))

    // strip everything in imgUrl after the last / for the filename
    offset := strings.LastIndex(imgUrl, `/`)
    filename := imgUrl[offset + 1:]
    fmt.Printf("%s\n", filename)

    response, err := http.Get(imgUrl)
    rosebud(err)
    image, err := ioutil.ReadAll(response.Body)
    rosebud(err)
    err = ioutil.WriteFile(filename, image, 0777)
    rosebud(err)
}


func main() {
    links := read_in_queue()

    for i := range links {
        if len(links[i]) == 0 {
            break
        }

        // loads a test page for the time being
        rootPage := load_url(links[i])

        // {1,5} in case of really huge galleries that I hope to never see.
        re := regexp.MustCompile(
            `http://g.e-hentai.org/s/[0-9a-f]{10}/[0-9]{6}-[0-9]{1,5}`)

        imageUrls := re.FindAllString(string(rootPage), -1)
        for q := range imageUrls {
            download_image(imageUrls[q])
        }

        fmt.Printf("%s", imageUrls)
    }
}
