package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

  "github.com/joho/godotenv"
)

func generateAlphabeticLabel(n int) string {
	var label string
	for n >= 0 {
		label = string('A'+(n%26)) + label
		n = n/26 - 1
	}
	return label
}

func openURLInConfiguredBrowser(browserPath string, url string) {
  err := exec.Command(browserPath, url).Run()
  if(err != nil) {
    fmt.Println("Could not open the configured browser: ", err)
		return
  }
}

func main() {
  err := godotenv.Load(".env")
  if err != nil {
    fmt.Println("Error loading .env file")
    return
  }

  workspace := os.Getenv("WORSKPACE")
	project := os.Getenv("PROJECT")
  browserPath := os.Getenv("BROWSER_PATH")
  projectPath := os.Getenv("PROJECT_PATH")
  
  if workspace == "" || project == "" || browserPath == "" || projectPath == "" {
    fmt.Println("Please set the following environment variables in the .env file: WORSKPACE, PROJECT, BROWSER_PATH, PROJECT_PATH")
    return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Search: ")
	prompt, _ := reader.ReadString('\n')
	prompt = strings.TrimSpace(prompt)

	cmd := exec.Command("git", "log", "--all", "--grep="+prompt)
  cmd.Dir = projectPath

	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Could not execute 'git log' commando: ", err)
		return
	}

	logOutput := string(output)
	commits := strings.Split(logOutput, "\n\n")

	commitRegex := regexp.MustCompile(`^commit (\w+)`)

	urlMap := make(map[string]string)
	labelCounter := 0

	for _, commit := range commits {
		if commit == "" {
			continue
		}

		matches := commitRegex.FindStringSubmatch(commit)
		if len(matches) > 1 {
			commitId := matches[1]
			url := fmt.Sprintf("https://%s/%s/%s", workspace, project, commitId)

			label := generateAlphabeticLabel(labelCounter)
			labelCounter++

			urlMap[label] = url
			open := fmt.Sprintf("URL (%s): %s\n ", label, url)
			fmt.Println(open)
			fmt.Println(commit)
			fmt.Println()
		}
	}

	fmt.Print("Open commit: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if url, exists := urlMap[choice]; exists {
		fmt.Println("Opening: ", url)
    openURLInConfiguredBrowser(browserPath, url)
		// openURL(url)
	} else {
		fmt.Println("No commit for shortcut selected")
	}
}
