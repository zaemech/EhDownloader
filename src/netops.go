package main

import (
    "io/ioutil"
    "os"
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


func download_image(imgUrl string) ([]byte, string) {
    // strip everything in imgUrl after the last / for the filename
    offset := strings.LastIndex(imgUrl, `/`)
    filename := imgUrl[offset + 1:]

    response, err := http.Get(imgUrl)
    rosebud(err)
    defer response.Body.Close()

    image, err := ioutil.ReadAll(response.Body)
    rosebud(err)
    return image, filename
}
