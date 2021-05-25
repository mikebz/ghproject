package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	ghproject "mikebz.com/ghproject/githandle"
)

func main() {
	godotenv.Load()
	testToken := os.Getenv("GITHUB_TOKEN")
	if testToken == "" {
		log.Fatal("GITHUB_TOKEN is empty, please set the environment variable to github token or create a .env file")
	}

	gh := ghproject.GitHandle{}
	ctx := context.Background()
	err := gh.Init(ctx, testToken)
	if err != nil {
		log.Fatal("Could not init GitHub handle with ", testToken,
			", Error: ", err)
	}

	issues, err := gh.SearchIssues("state:open repo:GoogleContainerTools/kpt milestone:\"v1.0 m3\"", ctx)
	if err != nil {
		log.Fatal("Could not get the test token ", err)
	}

	log.Println("Total issues: ", len(issues))

	for _, issue := range issues {
		assignee := issue.GetAssignee()

		// concat all the labels
		var builder strings.Builder
		for _, label := range issue.Labels {
			if builder.Len() > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(*label.Name)
		}

		log.Println(issue.GetNumber(), assignee.GetLogin(), ": ", builder.String())
	}
}
