package main

import (
    "testing"
)


type urlpair struct {
    url string
    clean string
}


var urlTest = []urlpair{
    // normal
    {"http://g.e-hentai.org/g/012345/abcde56789/",
    "http://g.e-hentai.org/g/012345/abcde56789/"},

    // normal with page number
    {"http://g.e-hentai.org/g/012345/abcde56789/?p=1",
    "http://g.e-hentai.org/g/012345/abcde56789/"},

    // normal with insanely large page number
    {"http://g.e-hentai.org/g/012345/abcde56789/?p=532187674",
    "http://g.e-hentai.org/g/012345/abcde56789/"},

    // missing page get request
    {"http://g.e-hentai.org/g/012345/abcde56789/?",
    "http://g.e-hentai.org/g/012345/abcde56789/"},

    // malformed page get request (missing =)
    {"http://g.e-hentai.org/g/012345/abcde56789/?p",
    "http://g.e-hentai.org/g/012345/abcde56789/"},

    // malformed page get request (missing page number)
    {"http://g.e-hentai.org/g/012345/abcde56789/?p=",
    "http://g.e-hentai.org/g/012345/abcde56789/"},

    // missing trailing '/'
    {"http://g.e-hentai.org/g/012345/abcde56789",
    "http://g.e-hentai.org/g/012345/abcde56789/"},

    // missing http
    {"g.e-hentai.org/g/012345/abcde56789/",
    "http://g.e-hentai.org/g/012345/abcde56789/"},

    // using https instead of http
    {"https://g.e-hentai.org/g/012345/abcde56789/",
    "http://g.e-hentai.org/g/012345/abcde56789/"},


    {"/012345/abcde56789/",
    "http://g.e-hentai.org/g/012345/abcde56789/"},


    {"012345/abcde56789/",
    "http://g.e-hentai.org/g/012345/abcde56789/"},


    {"012345/abcde56789",
    "http://g.e-hentai.org/g/012345/abcde56789/"},


    {"/012345/abcde56789",
    "http://g.e-hentai.org/g/012345/abcde56789/"},


    {"/012345/abcde56789/?p=1",
    "http://g.e-hentai.org/g/012345/abcde56789/"},
}


func Test_clean_url(t *testing.T) {
    for _, pair := range urlTest {
        result := clean_url(pair.url)
        if result != pair.clean {
            t.Error(result)
        }
    }
}
