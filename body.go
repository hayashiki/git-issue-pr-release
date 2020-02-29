package gipr

import (
	"log"
	"strconv"
	"strings"

	"github.com/google/go-github/github"

	"regexp"
)

func generateBody(Comparison *github.CommitsComparison) ([]int, []int, error) {
	mergedPRMsgExp := regexp.MustCompile(`#([0-9]+)`)
	var mergedPRNums []int
	var refIssueOrPRNums []int

	for _, c := range Comparison.Commits {

		log.Printf("commit is, %v", c.GetCommit().GetMessage())
		m := mergedPRMsgExp.FindSubmatch([]byte(c.GetCommit().GetMessage()))

		if len(m) > 1 {
			log.Printf("m0 is, %v", string(m[0]))
			log.Printf("m1 is, %v", string(m[1]))
			// log.Printf("m1 is, %v", m[1])
			n, err := strconv.Atoi(string(m[1]))
			if err != nil {
				log.Fatalf("error is %s", err)
				return nil, nil, nil
			}
			if strings.HasPrefix(c.GetCommit().GetMessage(), "Merge pull request") {
				mergedPRNums = append(mergedPRNums, n)
			} else {
				refIssueOrPRNums = append(refIssueOrPRNums, n)
			}
		}
	}
	log.Printf("mergedPRNums is, %v", mergedPRNums)
	log.Printf("refIssueNums is, %v", refIssueOrPRNums)
	return refIssueOrPRNums, mergedPRNums, nil
}

// *github.CommitsComparison{
// 	BaseCommit      *RepositoryCommit `json:"base_commit,omitempty"`
// 	MergeBaseCommit *RepositoryCommit `json:"merge_base_commit,omitempty"`

// 	// Head can be 'behind' or 'ahead'
// 	Status       *string `json:"status,omitempty"`
// 	AheadBy      *int    `json:"ahead_by,omitempty"`
// 	BehindBy     *int    `json:"behind_by,omitempty"`
// 	TotalCommits *int    `json:"total_commits,omitempty"`

// 	Commits []RepositoryCommit `json:"commits,omitempty"`

// 	Files []CommitFile `json:"files,omitempty"`

// 	HTMLURL      *string `json:"html_url,omitempty"`
// 	PermalinkURL *string `json:"permalink_url,omitempty"`
// 	DiffURL      *string `json:"diff_url,omitempty"`
// 	PatchURL     *string `json:"patch_url,omitempty"`
// 	URL          *string `json:"url,omitempty"` // API URL.
// }
