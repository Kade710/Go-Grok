package main

import (
    "fmt"
    "regexp"
)

// Minimal Grok struct
type Grok struct {
    patterns map[string]string
    regex    *regexp.Regexp
}

// Constructor
func NewGrok() *Grok {
    return &Grok{
        patterns: make(map[string]string),
    }
}

// Add custom patterns
func (g *Grok) AddPatterns(p map[string]string) {
    for k, v := range p {
        g.patterns[k] = v
    }
}

// Compile a pattern into a regex
func (g *Grok) Compile(pattern string) error {
    // Replace %{PATTERN} with actual regex for demonstration
    for name, pat := range g.patterns {
        pattern = regexp.MustCompile("%\\{"+name+"\\}").ReplaceAllString(pattern, pat)
    }
    r, err := regexp.Compile(pattern)
    if err != nil {
        return err
    }
    g.regex = r
    return nil
}

// Parse string using compiled regex
func (g *Grok) ParseString(input string) map[string]string {
    result := make(map[string]string)
    if g.regex == nil {
        return result
    }
    matches := g.regex.FindStringSubmatch(input)
    for i, name := range g.regex.SubexpNames() {
        if i != 0 && name != "" {
            result[name] = matches[i]
        }
    }
    return result
}

func main() {
    g := NewGrok()

    patterns := map[string]string{
        "NGINX_HOST":         `(?P<destination_ip>\d+\.\d+\.\d+\.\d+)(:(?P<destination_port>\d+))?`,
        "NGINX_NOTSEPARATOR": `[^\t ,:]+`,
    }
    g.AddPatterns(patterns)

    err := g.Compile("%{NGINX_HOST}")
    if err != nil {
        panic(err)
    }

    res := g.ParseString("127.0.0.1:1234")
    fmt.Println(res)
}

