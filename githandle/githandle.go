package ghproject

import (
	"context"
	"errors"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

const INIT_ERROR = "the the client needs to be initialized.  Call Init first"

type GitHandle struct {
	client *github.Client
}

// init function that calls the Github API
func (gh *GitHandle) Init(ctx context.Context, token string) error {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	gh.client = github.NewClient(tc)

	return nil
}

// assumes that you already have an initiated handle
// and enumerates the
func (gh *GitHandle) ListRepos(ctx context.Context) (repositories []*github.Repository, err error) {
	if gh.client == nil {
		return nil, errors.New(INIT_ERROR)
	}

	repositories, _, err = gh.client.Repositories.List(ctx, "", nil)

	return repositories, err
}

// call this to get the list of issues from github
// when there is an error the issue array is going to be nil.
func (gh *GitHandle) SearchIssues(q string, ctx context.Context) ([]*github.Issue, error) {
	if gh.client == nil {
		return nil, errors.New(INIT_ERROR)
	}
	opts := &github.SearchOptions{}
	result, _, err := gh.client.Search.Issues(ctx, q, opts)
	if err != nil {
		return nil, err
	}

	// TODO: handle the multi page return
	return result.Issues, nil
}
