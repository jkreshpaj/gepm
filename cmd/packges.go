package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// ? Package type contains all the package information
type Package struct {
	Name        string `json:"name"`
	Synopsis    string `json:"synopsis"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Package     string `json:"package"`
	ProjectURL  string `json:"projecturl"`
}

// ? Packages type contains search query and list of packages
type Packages struct {
	Query string    `json:"query"`
	Hits  []Package `json:"hits"`
}

// ? Search api for package and populate table
func searchPackages(q string, done chan bool) {
	res, err := http.Get("https://go-search.org/api?action=search&q=" + q)
	if err != nil {
		failed.Println("Error making search request")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		failed.Println("Failed to read response body")
	}
	json.Unmarshal([]byte(body), &packages)

	for p := 0; p < len(packages.Hits); p++ {
		i := strconv.Itoa(p + 1)
		packages.Hits[p].Name = "[" + i + "] " + packages.Hits[p].Name
	}

	fmt.Println("\n\u001b[1m\u001b[4m\u001b[7mSearch results for", os.Args[1], "\u001b[0m")
	populateTable(packages)
	done <- true
}

// ? Add pakages result to table
func populateTable(packages *Packages) {
	w, h := getTSize()
	box := drawBox(w, h)
	for i, row := range box {
		if i > 1 {
			if i-2 < len(box)-3 {
				if len(packages.Hits) > i-2 {
					writeToRow(row, packages.Hits[i-2])
				} else {
					writeToRow(row, Package{Name: "", Synopsis: "", Author: ""})
				}
			}
		} else if i == 1 {
			*row = setTitle(row)
		}
		fmt.Println(row.String())
	}
}
