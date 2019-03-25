// Main logic

package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
)

// All types are just clones of the JSON returned by a Tweet request call

type Media struct {
    Type                string  `json:"type"`
    MediaURLHttps       string  `json:"media_url_https"`
    ExpandedURL         string  `json:"expanded_url"`
}

type Entities struct {
    Media       []Media   `json:"media"`
}

type User struct {
    ScreenName  string  `json:"screen_name"`
}

type Tweet struct {
    CreatedAt           string          `json:"created_at"`
    DisplayTextRange    []int           `json:"display_text_range"`
    FullText            string          `json:"full_text"`
    Entities            Entities        `json:"extended_entities"`
    User                User            `json:"user"`
}

func (t *Tweet) CheckMedia() {

// The HTML display template only writes Media.MediaURLHttps, which is ideal
// for directly linking to images, but links to the thumbnail for any other
// media type. This procedure assigns Media.ExpandedURL to Media.MediaURLHTTPS
// for non-image media, so that it may be accessed properly.

    for i, _ := range t.Entities.Media {
        if t.Entities.Media[i].Type != "photo" {
            t.Entities.Media[i].MediaURLHttps = t.Entities.Media[i].ExpandedURL
        }
    }
}

func (t *Tweet) Print(w http.ResponseWriter) {

// Displays a tweet in plaintext format.

    x, y := t.DisplayTextRange[0], t.DisplayTextRange[1]
    fmt.Fprintln(w, t.User.ScreenName)
    fmt.Fprintln(w, t.CreatedAt)
    fmt.Fprintln(w, t.FullText[x:y])
    fmt.Fprintln(w, "--")
    for _, m := range t.Entities.Media {
        fmt.Fprintln(w, m.Type, m.MediaURLHttps)
    }
}

func (t *Tweet) PrintJSON(w http.ResponseWriter) {

// Displays a tweet in JSON format.    
    s, err := json.Marshal(t)
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Fprintln(w, string(s))
}

func (t *Tweet) PrintHTML(w http.ResponseWriter) {

// Displays a tweet as HTML.

    x, y := t.DisplayTextRange[0], t.DisplayTextRange[1]
    t.FullText = t.FullText[x:y]
    TweetTemplate.Execute(w, t)
}

func serve(w http.ResponseWriter, r *http.Request) {

// This is the solitary HTTP listener function. When provided no arguments,
// it prints the form HTMLSearch as a web interface. It prints the tweet
// pointed to by url= as format= otherwise. The default format is plaintext,
// for simple command-line parsing, but the default format provided by the 
// form is HTML.

    u, format := "", ""
    t := Tweet{}
    k, _ := r.URL.Query()["url"]
    if len(k) > 0 {
        u = k[0]
    }
    k, _ = r.URL.Query()["format"]
    if len(k) > 0 {
        format = strings.ToLower(k[0])
    }
    if u == "" {
        fmt.Fprintf(w, SearchHTML)
        return
    }
    // Only tweets as "https://twitter.com/user/status/id" are honored at the
    // moment. I'm not sure if there are more exotic arrangements that should
    // be accounted for.
    id := strings.Split(u, "/")[5]
    client := &http.Client{}
    req, _ := http.NewRequest("GET", T_URL + id, nil)
    req.Header.Add("Authorization", TOKEN)
    res, _ := client.Do(req)
    body, _ := ioutil.ReadAll(res.Body)
    json.Unmarshal(body, &t)
    t.CheckMedia()
    if format == "json" {
        t.PrintJSON(w)
    } else if format == "html" {
        t.PrintHTML(w)
    } else {
        t.Print(w)
    }
}

func main() {
    
// Check for the solitary port argument, then listen on it. Nothing fancy.

    if len(os.Args) != 2 {
        log.Fatal("Usage: bitter port")
    }
    http.HandleFunc("/", serve)
    err := http.ListenAndServe(":" + os.Args[1], nil)
    if err != nil {
        log.Fatal(err)
    }
}
