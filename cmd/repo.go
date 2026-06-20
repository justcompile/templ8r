/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"templ8r/internal/cache"

	"templ8r/internal/types"

	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use: "repo",
}

var addCmd = &cobra.Command{
	Use:  "add <name> <url>",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		repo := &types.Repo{
			Name:   args[0],
			URL:    args[1],
			Branch: cmd.Flag("branch").Value.String(),
			Path:   cmd.Flag("path").Value.String(),
		}

		if err := cache.RepoController.AddRepository(repo); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println("Repository added successfully.")
	},
}

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use: "rm",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rm called")
	},
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		for _, repo := range cache.RepoController.ListRepositories() {
			fmt.Println(repo)
		}
	},
}

func init() {
	addCmd.Flags().StringP("branch", "b", "main", "Repository branch to use (default is 'main')")
	addCmd.Flags().StringP("path", "p", "", "subdirectory path within the repository to use (default is root)")
	repoCmd.AddCommand(addCmd)

	repoCmd.AddCommand(listCmd)

	repoCmd.AddCommand(rmCmd)

	rootCmd.AddCommand(repoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
