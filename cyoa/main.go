package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"unicode/utf8"
)

func main() {
	
	story := parseJSON("gopher.json")
	h := cyoaHandler{story: story}
	// fmt.Println(story)
	
	// minimal example
	// tmpl := template.Must(template.ParseFiles("cyoa_layout.html"))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl.Execute(w , story["intro"])
	// })
	http.ListenAndServe(":3000", h)
}

func parseJSON(filename string) map[string]storyStruct {
	// 1. open file
	jsonfile, err := os.Open(filename)
	printError(err)
	fmt.Println("Successfully opened json file!")
	defer jsonfile.Close()
	// 2. read file
	bytevalue, _ := ioutil.ReadAll(jsonfile)
	// 3. unmarshal json
	var out map[string]storyStruct
	err2 := json.Unmarshal(bytevalue, &out)
	printError(err2)
	// 4. alternative way for 2. and 3., using a decoder
	// decJSON := json.NewDecoder(jsonfile)
	// err3 := decJSON.Decode(&out)
	// printError(err3)
	return out
}

type storyStruct struct {
	Title string 		`json:"title"`
	Story []string 		`json:"story"`
	Options []option 	`json:"options"`
}

type option struct {
	Text string	`json:"text"`
	Arc string	`json:"arc"`
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type cyoaHandler struct {
	story map[string]storyStruct
}

func (h cyoaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := trimFirstRune(r.URL.Path)
	tmpl := template.Must(template.ParseFiles("cyoa_layout.html")) // Must() panics if call to ParseFiles is non-nil
	if page, ok := h.story[path]; ok {
		tmpl.Execute(w , page)
	} else if path == "" {
		page := h.story["intro"]
		tmpl.Execute(w , page)
	} else if path == "styles.css" {
		http.ServeFile(w, r, "styles.css")
	} else {
		http.NotFound(w, r)
	}
}

// source: https://stackoverflow.com/questions/48798588/how-do-you-remove-the-first-character-of-a-string/48798712
func trimFirstRune(s string) string {
    _, i := utf8.DecodeRuneInString(s)
    return s[i:]
}