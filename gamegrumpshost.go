package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var html string

func requestHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, html)
}

func main() {
	defer func() {
		fmt.Println(recover())
	}()
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	rawhtml, err := ioutil.ReadFile(strings.Join([]string{dir, "gamegramps.html"}, "/"))
	if err != nil {
		panic(err.Error())
	}
	html = string(rawhtml[:])
	fmt.Print("Hosting @ localhost:8080 gogo game groomples <3")
	http.HandleFunc("/gamegramps", requestHandler)
	http.ListenAndServe(":8080", nil)
}
