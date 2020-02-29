package gh

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

//go:generate mockgen -source github.go -destination mock_github/github.go
type Github interface {
	CompareCommits(ctx context.Context, owner, repo, base, head string) (*github.CommitsComparison, error)
	GetIsssues(ctx context.Context, owner, repo string, issueNumber int) (*github.Issue, error)
	GetPullRequests(ctx context.Context, owner, repo string, prNumber int) (*github.PullRequest, error)
	CreatePullRequest(ctx context.Context, owner, repo string, npr *github.NewPullRequest) (*github.PullRequest, error)
}

// Github represents Github client
type Client struct {
	cli *github.Client
}

// NewGitLabClient returns new Github interface
func NewGithubClient(ctx context.Context, token string) *Client {
	return &Client{
		github.NewClient(
			oauth2.NewClient(
				ctx, oauth2.StaticTokenSource(
					&oauth2.Token{AccessToken: token},
				),
			),
		),
	}
}

func (c *Client) CompareCommits(ctx context.Context, owner, repo, base, head string) (*github.CommitsComparison, error) {
	comp, _, err := c.cli.Repositories.CompareCommits(ctx, owner, repo, base, head)
	if err != nil {
		return nil, err
	}

	return comp, nil
}

func (c *Client) GetIsssues(ctx context.Context, owner, repo string, issueNumber int) (*github.Issue, error) {
	issue, _, err := c.cli.Issues.Get(ctx, owner, repo, issueNumber)
	if err != nil {
		return nil, err
	}

	return issue, err
}

func (c *Client) GetPullRequests(ctx context.Context, owner, repo string, prNumber int) (*github.PullRequest, error) {
	pr, _, err := c.cli.PullRequests.Get(ctx, owner, repo, prNumber)
	if err != nil {
		return nil, err
	}

	return pr, err
}

func (c *Client) CreatePullRequest(ctx context.Context, owner, repo string, npr *github.NewPullRequest) (*github.PullRequest, error) {
	pr, _, err := c.cli.PullRequests.Create(ctx, owner, repo, npr)
	if err != nil {
		return nil, err
	}

	return pr, err
}
