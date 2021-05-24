package ghproject

import (
	"context"
	"errors"

	"log"

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
	opts.Page = 1

	var allIssues []*github.Issue

	for more := true; more; {
		result, response, err := gh.client.Search.Issues(ctx, q, opts)
		if err != nil {
			log.Printf("Received an error from search %v", err)
			return nil, err
		}

		allIssues = append(allIssues, result.Issues...)
		log.Printf("Length of allIssues is %d", len(allIssues))
		log.Printf("Response.LastPage is %d and opts.Page is %d", response.LastPage, opts.Page)

		// looking at the result of the API it looks like they set the LastPage
		// to zero when you are on the last page.
		if response.LastPage == 0 {
			more = false
		} else {
			opts.Page++
		}
	}

	return allIssues, nil
}

// call this to get a list of milestones that are open
func (gh *GitHandle) ListMilestones(owner string, repo string, ctx context.Context) ([]*github.Milestone, error) {
	issuesService := gh.client.Issues
	// TODO: deal with multiple pages (eventually)
	milestones, _, err := issuesService.ListMilestones(ctx, owner, repo, nil)
	return milestones, err
}
