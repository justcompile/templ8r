package cache

import (
	"os"
	"path/filepath"
	"templ8r/pkg/utils"
)

func getCacheDirPath() string {
	home := utils.Must(os.UserHomeDir())
	return filepath.Join(home, ".config", "templ8r")
}

func getRepositoriesFilePath() string {
	return filepath.Join(getCacheDirPath(), "repositories.yaml")
}

func Setup() error {
	dir := getCacheDirPath()

	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	return initializeRepositories()
}
