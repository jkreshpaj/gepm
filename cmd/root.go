package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type Package struct {
	Name        string `json:"name"`
	Synopsis    string `json:"synopsis"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Package     string `json:"package"`
	ProjectURL  string `json:"projecturl"`
}

type Packages struct {
	Query string    `json:"query"`
	Hits  []Package `json:"hits"`
}

var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "TEST IS TEST",
	Long:  "TEST TEST TEST IS TEST",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get("https://go-search.org/api?action=search&q=" + os.Args[1])
		if err != nil {
			log.Println("ERROR MAKING REQUEST")
		}
		pcks := new(Packages)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("ERROR READING BODY")
		}
		json.Unmarshal([]byte(body), &pcks)
		for p := 0; p < len(pcks.Hits); p++ {
			i := strconv.Itoa(p + 1)
			pcks.Hits[p].Name = "[" + i + "] " + pcks.Hits[p].Name
		}
		fmt.Println("\u001b[1m\u001b[4m\u001b[7mSearch results for", os.Args[1], "\u001b[0m")
		w, h := getTSize()
		boxArr := drawBox(w, h)
		for i, row := range boxArr {
			if i > 1 && i < len(pcks.Hits) {
				if i-2 < len(boxArr)-3 {
					writeToRow(row, pcks.Hits[i-2])
				}
			} else if i == 1 {
				*row = setTitle(row)
			}
			fmt.Println(row.String())
		}
		fmt.Print("➡ Enter number of package to be installed: ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		pn, _ := strconv.Atoi(input.Text())
		fmt.Println("➡ Installing package", pcks.Hits[pn-1].Name, "from", pcks.Hits[pn-1].Package)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
