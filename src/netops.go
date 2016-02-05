package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "regexp"
    "net/http"
    "strings"
)

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
