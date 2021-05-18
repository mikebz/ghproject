package ghproject

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// test the init function
func TestInit(t *testing.T) {
	testToken := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	gh := GitHandle{}
	err := gh.Init(ctx, testToken)
	assert.NoError(t, err)
}

// test repositories
// also relies on the Init to be working
func TestListRepos(t *testing.T) {
	testToken := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	gh := GitHandle{}
	err := gh.Init(ctx, testToken)
	assert.Nil(t, err)

	repos, err := gh.ListRepos(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, repos)
}

// search for open issues in the kpt repo.  Should return some results.
func TestSearchIssues(t *testing.T) {
	testToken := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	gh := GitHandle{}
	err := gh.Init(ctx, testToken)
	assert.Nil(t, err)

	issues, err := gh.SearchIssues("state:open repo:GoogleContainerTools/kpt", ctx)
	assert.NoError(t, err)
	assert.NotNil(t, issues)
	assert.Greater(t, len(issues), 0)
}
