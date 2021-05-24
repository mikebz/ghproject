package ghproject

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var gh GitHandle
var ctx context.Context

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	testToken := os.Getenv("GITHUB_TOKEN")
	gh = GitHandle{}
	ctx = context.Background()
	err = gh.Init(ctx, testToken)
	if err != nil {
		log.Fatal("Could not init GitHub handle with ", testToken,
			", Error: ", err)
	}

	os.Exit(m.Run())
}

// test repositories
// also relies on the Init to be working
func TestListRepos(t *testing.T) {
	repos, err := gh.ListRepos(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, repos)
}

// search for open issues in the kpt repo.  Should return some results.
func TestSearchIssues(t *testing.T) {
	issues, err := gh.SearchIssues("state:open repo:GoogleContainerTools/kpt", ctx)
	assert.NoError(t, err)
	assert.NotNil(t, issues)
	assert.Greater(t, len(issues), 0)
}

func TestSearchIssuesMilestone(t *testing.T) {
	issues, err := gh.SearchIssues("state:open repo:GoogleContainerTools/kpt milestone:\"v1.0 m3\"", ctx)
	assert.NoError(t, err)
	assert.NotNil(t, issues)
	assert.Greater(t, len(issues), 0)
}

func TestListMilestones(t *testing.T) {
	milestones, err := gh.ListMilestones("GoogleContainerTools", "kpt", ctx)
	assert.NoError(t, err)
	assert.NotNil(t, milestones)
	assert.Greater(t, len(milestones), 0)
}
