package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// ? Ask for package number to be installed
func makePrompt() {
	installed := make(chan bool)
	packageNumber := makeQuestion()
	progress.Prefix = "➡ Installing package " + packages.Hits[packageNumber].Name + " from " + packages.Hits[packageNumber].Package + " "

	progress.Start()
	go makeInstall(packages, packageNumber, installed)
	<-installed
	progress.Stop()

	success.Println("➡ Successfully installed", packages.Hits[packageNumber].Name)
}

// ? Go get package
func makeInstall(packages *Packages, packageNumber int, installed chan bool) {
	cmd := exec.Command("go", "get", packages.Hits[packageNumber].Package)
	cmd.Run()

	saveToFile(packages, packageNumber)
	installed <- true
}

// ? Make package number question
func makeQuestion() int {
	fmt.Print("➡ Enter number of package to be installed: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	packageNumber, _ := strconv.Atoi(input.Text())

	// Check if input is int if not remake question
	if _, err := strconv.Atoi(input.Text()); err != nil {
		failed.Println("Please type a package number or Ctrl+C to cancel")
		return makeQuestion()
	}

	return packageNumber - 1
}

// ? Save installed package to file
func saveToFile(packages *Packages, packageNumber int) {
	newPackage := make(map[string]interface{})

	// Open packages.json
	file, err := os.OpenFile("packages.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		failed.Println("Error opening packages.json")
	}
	defer file.Close()

	// Read packages.json
	depsByte, err := ioutil.ReadFile("packages.json")
	if err != nil {
		failed.Println("Error reading packages.json")
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(depsByte), &result)
	if result == nil {
		newPackage[strings.Split(packages.Hits[packageNumber].Name, " ")[1]] = packages.Hits[packageNumber].Package
		result = newPackage
	} else {
		result[strings.Split(packages.Hits[packageNumber].Name, " ")[1]] = packages.Hits[packageNumber].Package
	}

	// Write packages.json
	deps, _ := json.MarshalIndent(result, "", "\t")
	err = ioutil.WriteFile("packages.json", deps, 0644)
	if err != nil {
		failed.Println("Error writing to packages.json")
	}
}

func makeInstallFromFile() {
	file, err := os.OpenFile("packages.json", os.O_RDONLY, 0600)
	if err != nil && file == nil {
		failed.Println("Unable to find packages.json on current directory")
	} else {
		packagesByte, err := ioutil.ReadFile("packages.json")
		if err != nil {
			failed.Println("Error reading packages.json")
		}
		var result map[string]interface{}
		json.Unmarshal([]byte(packagesByte), &result)

		for key := range result {
			fmt.Println("➡ Installing package", key, "from", result[key])
			url, _ := result[key].(string)
			cmd := exec.Command("go", "get", url)
			cmd.Run()
			success.Println("➡ Successfully installed", key)
		}
	}
	defer file.Close()
}
