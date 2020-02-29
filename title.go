package gipr

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func generateTitle() (string, error) {
	// title = "Merge develop->staging"
	content, err := ioutil.ReadFile(".github/TITLE_TEMPLATE.md")
	if err != nil {
		return "", errors.New(fmt.Sprintf("could not generate PR: %v", err))
	}
	return string(content), nil
}
