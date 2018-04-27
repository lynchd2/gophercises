package main 


import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

var story map[string]interface{}


type Chapter struct {
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}



func convertJson() map[string]interface{} {
	file, err := ioutil.ReadFile("./story.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(file, &story)
	return story
	
}


func introHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}


func main() {
	json := convertJson()
	fmt.Println(json)
	http.HandleFunc("/", introHandler)
	http.ListenAndServe(":8080", nil)

}