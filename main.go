package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/akagami-harsh/Experience/Jaeger"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	username = "akagami-harsh"
)

// RepoConfig holds configuration for a specific repository
type RepoConfig struct {
	Owner    string
	Repo     string
	DataFile string
	MDFile   string
}

// Template for creating a new data file
const dataTemplate = `package {{.Package}}

const Data = ` + "`" + `
# {{.Repo}} - My Contribution Journey

## Little about me

- Name: Harshvir Potpose
- Email: <hpotpose62@gmail.com>
- GitHub Username: [akagami-harsh](https://github.com/akagami-harsh)

## About the project

{{.Repo}} is an open-source project. 
<!-- Add more details about the project here -->

## Technical Contributions

### [{{.Repo}}](https://github.com/{{.Owner}}/{{.Repo}})

View all pull requests by me at a glance: [VIEW ALL PULL REQUESTS](https://github.com/{{.Owner}}/{{.Repo}}/pulls?q=is%3Apr+author%3A{{.Username}}+is%3Aclosed)
` + "`"

func main() {
	// Parse command line arguments
	var owner, repo string
	flag.StringVar(&owner, "owner", "", "GitHub repository owner")
	flag.StringVar(&repo, "repo", "", "GitHub repository name")
	flag.Parse()

	if owner == "" || repo == "" {
		// If no arguments are provided, use Jaeger as default
		owner = "jaegertracing"
		repo = "jaeger"
		fmt.Println("No repository specified. Using default: jaegertracing/jaeger")
		fmt.Println("Usage: go run main.go -owner=<owner> -repo=<repo>")
	}

	// Create repository directories and configuration
	repoConfig := setupRepository(owner, repo)

	// Fetch pull requests
	prs := fetchPullRequests(owner, repo)

	// Generate README
	generateReadme(repoConfig, prs)

	fmt.Printf("Markdown file '%s' created.\n", repoConfig.MDFile)

	// Update main README to include the new repository
	updateMainReadme(owner, repo)
}

func setupRepository(owner, repo string) RepoConfig {
	// Create a directory for the repository if it doesn't exist
	dirName := strings.Title(repo)
	dirPath := filepath.Join(".", dirName)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			os.Exit(1)
		}
	}

	// Configure paths for the repository
	dataFile := filepath.Join(dirPath, "readmeData.go")
	mdFile := filepath.Join(dirPath, "README.md")

	// Create data file if it doesn't exist
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		createDataFile(owner, repo, dirName, dataFile)
	}

	return RepoConfig{
		Owner:    owner,
		Repo:     repo,
		DataFile: dataFile,
		MDFile:   mdFile,
	}
}

func createDataFile(owner, repo, packageName, dataFile string) {
	// Create template data
	data := struct {
		Owner    string
		Repo     string
		Package  string
		Username string
	}{
		Owner:    owner,
		Repo:     repo,
		Package:  packageName,
		Username: username,
	}

	// Create template
	tmpl, err := template.New("data").Parse(dataTemplate)
	if err != nil {
		fmt.Printf("Error creating template: %v\n", err)
		os.Exit(1)
	}

	// Create file
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Printf("Error creating data file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Execute template
	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created data file: %s\n", dataFile)
}

func fetchPullRequests(owner, repo string) []*github.PullRequest {
	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("Set the GITHUB_TOKEN environment variable.")
		os.Exit(1)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	prs := make([]*github.PullRequest, 0)
	for i := 1; i <= 30; i++ {
		opts := &github.PullRequestListOptions{
			State:       "closed",
			Head:        username,
			ListOptions: github.ListOptions{PerPage: 100, Page: i},
		}
		pr, _, err := client.PullRequests.List(ctx, owner, repo, opts)
		if err != nil {
			fmt.Printf("Error fetching pull requests: %v\n", err)
			os.Exit(1)
		}
		prs = append(prs, pr...)
		fmt.Println("Page:", i)

		// Break if we got less than the requested number of PRs
		if len(pr) < opts.PerPage {
			break
		}
	}

	filteredPRs := make([]*github.PullRequest, 0)
	for _, pull := range prs {
		if *pull.User.Login == username {
			filteredPRs = append(filteredPRs, pull)
		}
	}

	fmt.Printf("Found %d pull requests for %s/%s\n", len(filteredPRs), owner, repo)
	return filteredPRs
}

func generateReadme(config RepoConfig, prs []*github.PullRequest) {
	// Import package dynamically based on repo
	var data string

	// For now, we only have the Jaeger package hardcoded
	// In a more advanced version, we would dynamically import packages
	if strings.ToLower(config.Repo) == "jaeger" {
		data = Jaeger.Data
	} else {
		// For other repos, read from the data file (this is a simplification)
		// In a real implementation, you would dynamically import the package
		content, err := os.ReadFile(config.DataFile)
		if err != nil {
			fmt.Printf("Error reading data file: %v\n", err)
			os.Exit(1)
		}

		// Extract the Data constant value
		parts := strings.Split(string(content), "const Data = `")
		if len(parts) < 2 {
			fmt.Println("Invalid data file format")
			os.Exit(1)
		}

		dataPart := parts[1]
		data = dataPart[:strings.LastIndex(dataPart, "`")]
	}

	var sb strings.Builder
	sb.WriteString(data)
	sb.WriteString("\n\n")

	sb.WriteString("| Date Created | Title | Pull Request Link |\n")
	sb.WriteString("| ------------ | ----- | ----------------- |\n")

	for _, pr := range prs {
		date := pr.CreatedAt.Format(time.DateOnly)
		title := strings.ReplaceAll(*pr.Title, "|", "\\|")
		url := *pr.HTMLURL
		sb.WriteString(fmt.Sprintf("| %s | %s | [PR link](%s) |\n", date, title, url))
	}

	err := os.WriteFile(config.MDFile, []byte(sb.String()), 0644)
	if err != nil {
		fmt.Printf("Error writing markdown file: %v\n", err)
		os.Exit(1)
	}
}

func updateMainReadme(owner, repo string) {
	// Read the current README
	readmePath := "./README.md"
	content, err := os.ReadFile(readmePath)
	if err != nil {
		fmt.Printf("Error reading main README: %v\n", err)
		return
	}

	readmeContent := string(content)

	// Check if the repository is already in the README
	repoSection := fmt.Sprintf("[%s - %s]", owner, repo)
	if strings.Contains(readmeContent, repoSection) {
		fmt.Println("Repository already in main README")
		return
	}

	// Find the open source contributions section
	contributionsSection := "- Open Source Contributions"
	contributionsIndex := strings.Index(readmeContent, contributionsSection)

	if contributionsIndex == -1 {
		fmt.Println("Could not find 'Open Source Contributions' section in README")
		return
	}

	// Find the line with "[Add more projects as needed]"
	projectsLine := "  - [Add more projects as needed]"
	projectsIndex := strings.Index(readmeContent, projectsLine)

	if projectsIndex == -1 {
		fmt.Println("Could not find placeholder for adding more projects")
		return
	}

	// Add the new repository entry before the placeholder
	dirName := strings.Title(repo)
	newRepoEntry := fmt.Sprintf("  - [%s - %s](https://github.com/akagami-harsh/Experience/blob/main/%s/README.md)\n", owner, repo, dirName)

	updatedContent := readmeContent[:projectsIndex] + newRepoEntry + readmeContent[projectsIndex:]

	// Write the updated content ack to the file
	err = os.WriteFile(readmePath, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Printf("Error updating main README: %v\n", err)
		return
	}

	fmt.Println("Updated main README with new repository")
}
