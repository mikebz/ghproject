package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	githandle "mikebz.com/ghproject/githandle"
	stats "mikebz.com/ghproject/stats"
)

func main() {
	godotenv.Load()
	testToken := os.Getenv("GITHUB_TOKEN")
	if testToken == "" {
		log.Fatal("GITHUB_TOKEN is empty, please set the environment variable to github token or create a .env file")
	}

	gh := githandle.GitHandle{}
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

	allData := make([]stats.IssueData, len(issues))

	for i, issue := range issues {
		issueData := stats.IssueData{}
		issueData.Assignee = issue.GetAssignee().GetLogin()

		issueData.Labels = make([]string, len(issue.Labels))
		for i, label := range issue.Labels {
			issueData.Labels[i] = label.GetName()
		}

		issueData.Id = issue.GetNumber()
		allData[i] = issueData
	}

	workload := stats.WorkloadByUser(allData)

	for user, days := range workload {
		fmt.Println("user: ", user, ", days: ", days)
	}
}
