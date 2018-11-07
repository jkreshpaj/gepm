package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
)

var (
	s        = spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	p        = spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	packages = new(Packages)
)

//Package type contains all the package information
type Package struct {
	Name        string `json:"name"`
	Synopsis    string `json:"synopsis"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Package     string `json:"package"`
	ProjectURL  string `json:"projecturl"`
}

//Packages type contains search query and list of packages
type Packages struct {
	Query string    `json:"query"`
	Hits  []Package `json:"hits"`
}

func searchPackages(q string, done chan bool) {
	res, err := http.Get("https://go-search.org/api?action=search&q=" + q)
	if err != nil {
		log.Println("ERROR MAKING REQUEST")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERROR READING BODY")
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

func populateTable(pcks *Packages) {
	w, h := getTSize()
	boxArr := drawBox(w, h)
	for i, row := range boxArr {
		if i > 1 {
			if i-2 < len(boxArr)-3 {
				if len(pcks.Hits) > i-2 {
					writeToRow(row, pcks.Hits[i-2])
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

func makePrompt() {
	installed := make(chan bool)
	fmt.Print("➡ Enter number of package to be installed: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	pn, _ := strconv.Atoi(input.Text())
	p.Prefix = "➡ Installing package " + packages.Hits[pn-1].Name + " from " + packages.Hits[pn-1].Package + " "
	p.Start()
	go makeInstall(packages, pn, installed)
	<-installed
	p.Stop()
	fmt.Println("➡ Successfully installed", packages.Hits[pn-1].Name)
}

func makeInstall(pcks *Packages, pn int, installed chan bool) {
	cmd := exec.Command("go", "get", pcks.Hits[pn-1].Package)
	cmd.Run()
	installed <- true
}
