package types

import (
	"io"
	"os"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
)

type Repo struct {
	// Name is the name of the repository.
	Name string `json:"name" yaml:"name"`

	// URL is the URL of the repository.
	URL string `json:"url" yaml:"url"`

	// Branch is the branch of the repository to use.
	Branch string `json:"branch" yaml:"branch"`

	// Path is the path within the repository to use.
	Path string `json:"path" yaml:"path"`
}

func (r *Repo) Clone(toDir string, withProgress bool) error {
	var progress io.Writer = nil
	if withProgress {
		progress = os.Stdout
	}

	if _, err := git.PlainClone(toDir, &git.CloneOptions{
		URL:           r.URL,
		ReferenceName: plumbing.NewBranchReferenceName(r.Branch),
		SingleBranch:  true,
		Progress:      progress,
	}); err != nil {
		return err
	}

	return nil
}
