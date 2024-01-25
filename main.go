package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

const (
	owner    = "jaegertracing"
	repo     = "jaeger"
	username = "akagami-harsh"
)

func main() {
	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("Set the GITHUB_TOKEN environment variable.")
		return
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Fetch pull requests
	opts := &github.PullRequestListOptions{
		State:       "closed",
		Head:        username,
		ListOptions: github.ListOptions{PerPage: 100},
	}
	prs, _, err := client.PullRequests.List(ctx, owner, repo, opts)
	if err != nil {
		fmt.Printf("Error fetching pull requests: %v\n", err)
		return
	}

	filteredPulls := make([]*github.PullRequest, 0)
	for _, pull := range prs {
		if *pull.User.Login == username {
			filteredPulls = append(filteredPulls, pull)
		}
	}

	// Create markdown table
	var sb strings.Builder
	sb.WriteString("| Date Created | Title | Pull Request Link |\n")
	sb.WriteString("| ------------ | ----- | ----------------- |\n")

	for _, pr := range filteredPulls {
		date := pr.CreatedAt.Format(time.DateOnly)
		title := strings.ReplaceAll(*pr.Title, "|", "\\|")
		url := *pr.HTMLURL
		sb.WriteString(fmt.Sprintf("| %s | %s | [PR link](%s) |\n", date, title, title, url))
	}

	mdFilename := "./Jaeger/contributions.md"
	err = os.WriteFile(mdFilename, []byte(sb.String()), 0644)
	if err != nil {
		fmt.Printf("Error writing markdown file: %v\n", err)
		return
	}

	fmt.Printf("Markdown file '%s' created.\n", mdFilename)
}
