package main

import (
    "fmt"
)

func main() {
    g := grok.New()

    patterns := map[string]string{
        "NGINX_HOST":         `(?:%{IP:destination.ip}|%{NGINX_NOTSEPARATOR:destination.domain})(:%{NUMBER:destination.port:int})?`,
        "NGINX_NOTSEPARATOR": `"[^\t ,:]+"`,
    }

    g.AddPatterns(patterns)
    err := g.Compile("%{NGINX_HOST}", true)
    if err != nil {
        panic(err)
    }

    res, err := g.ParseTypedString("127.0.0.1:1234")
    if err != nil {
        panic(err)
    }

    fmt.Println(res)
}

