package gipr

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/google/go-github/github"
	"github.com/hayashiki/git-issue-pr-release/gh"
	"github.com/hayashiki/git-issue-pr-release/git"
)

func Run(token, head, base string) {
	ctx := context.Background()
	originURL, err := git.GetRemoteOrigin(ctx)
	// originURL = []byte("git@github.com:watashino-okyoushitsu/manatea-server.git")

	if err != nil {
		log.Fatalf("error is %s", err)
		return
	}
	log.Printf("url is %s", originURL)

	exp := regexp.MustCompile(`git@github\.com:(?P<owner>.+)/(?P<repo>.+)`)
	match := exp.FindSubmatch(originURL)
	originNames := make(map[string]string)

	for i, name := range exp.SubexpNames() {
		log.Printf("name is, %s", name)
		if i != 0 && name != "" {
			originNames[name] = string(match[i])
		}
	}
	log.Printf("originNames is, %v", originNames)
	// originNames is, map[owner:hayashiki repo:git-issue-release.git]
	owner, repo := originNames["owner"], strings.Split(originNames["repo"], ".git")[0]
	log.Printf("owner is, %v", owner)
	// own
	log.Printf("repo is, %v", repo)
	// git-issue-release

	client := gh.NewGithubClient(ctx, token)
	// client := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *token})))

	comp, err := client.CompareCommits(ctx, owner, repo, base, head)

	if err != nil {
		log.Fatalf("error is %s", err)
		return
		// return "", errors.New(fmt.Sprintf("could not genereate PR body: %+v", err))
	}

	// log.Printf("comp => %s", comp)

	// mergedPRMsgExp := regexp.MustCompile(`^Merge\spull\srequest\s#([0-9]+).+`)

	title, err := generateTitle()
	if err != nil {
		log.Fatalf("error is %s", err)
		return
	}

	var issueBody = "# Ref Issue List\n"
	var prBody = "# Ref Pull Request List\n"
	var mergedPrBody = "# Merged Pull Request List\n"

	refIssueOrPRNums, mergedPRNums, err := generateBody(comp)

	for _, v := range refIssueOrPRNums {
		issue, err := client.GetIsssues(ctx, owner, repo, v)
		if err != nil {
			log.Fatalf("error is %s", err)
			return
		}
		if issue.IsPullRequest() {
			prBody += fmt.Sprintf("- [ ] [#%d](%s) %s created by @%s\n", v, issue.GetHTMLURL(), issue.GetTitle(), issue.GetUser().GetLogin())
		} else {
			issueBody += fmt.Sprintf("- [ ] [#%d](%s) %s created by @%s\n", v, issue.GetHTMLURL(), issue.GetTitle(), issue.GetUser().GetLogin())
		}
	}

	for _, v := range mergedPRNums {
		pr, err := client.GetPullRequests(ctx, owner, repo, v)
		if err != nil {
			log.Fatalf("error is %s", err)
			return
		}
		mergedPrBody += fmt.Sprintf("- [ ] [#%d](%s) %s created by @%s\n", v, pr.GetHTMLURL(), pr.GetTitle(), pr.GetUser().GetLogin())
	}

	body := issueBody + prBody + mergedPrBody
	log.Printf("body is, %v", body)
	return

	pr, err := client.CreatePullRequest(ctx, owner, repo, &github.NewPullRequest{
		Title:               github.String(title),
		Head:                github.String(head),
		Base:                github.String(base),
		Body:                github.String(body),
		MaintainerCanModify: github.Bool(true),
	})

	if err != nil {
		log.Fatalf("error is %s", err)
		return
	}

	fmt.Printf("PR created: %s\n", pr.GetHTMLURL())

	return
}
