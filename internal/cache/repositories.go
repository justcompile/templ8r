package cache

import (
	"errors"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"templ8r/internal/types"

	"go.yaml.in/yaml/v3"
)

type repositoryController struct {
	// Repositories is a map of repository names to their corresponding Repo structs.
	Repositories map[string]types.Repo `json:"repositories" yaml:"repositories"`
}

func (rc *repositoryController) AddRepository(repo *types.Repo) error {
	if _, exists := rc.Repositories[repo.Name]; exists {
		return errors.New("repository with this name already exists")
	}

	repoPath := filepath.Join(getCacheDirPath(), repo.Name)

	if err := repo.Clone(repoPath, true); err != nil {
		return err
	}

	rc.Repositories[repo.Name] = *repo
	return rc.flushToFile()
}

func (rc *repositoryController) ListRepositories() []types.Repo {
	return slices.Collect(maps.Values(rc.Repositories))
}

func (rc *repositoryController) flushToFile() error {
	f, err := os.Create(getRepositoriesFilePath())
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := yaml.NewEncoder(f)
	defer encoder.Close()

	return encoder.Encode(rc.Repositories)
}

var RepoController *repositoryController

func initializeRepositories() error {
	if RepoController != nil {
		return errors.New("repository controller already initialized")
	}

	f, err := os.Open(getRepositoriesFilePath())
	if err != nil {
		if os.IsNotExist(err) {
			// If the file doesn't exist, initialize an empty controller.
			RepoController = &repositoryController{
				Repositories: make(map[string]types.Repo),
			}
			return nil
		}
		return err
	}
	defer f.Close()

	var repos map[string]types.Repo

	if err := yaml.NewDecoder(f).Decode(&repos); err != nil {
		return err
	}

	RepoController = &repositoryController{
		Repositories: repos,
	}

	return nil
}
