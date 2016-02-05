package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "bufio"
    "os"
    "regexp"
    "net/http"
    "strings"
    "strconv"
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

/*
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
*/


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
