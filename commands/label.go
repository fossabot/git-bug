package commands

import (
	"errors"
	"fmt"

	"github.com/MichaelMure/git-bug/cache"
	"github.com/spf13/cobra"
)

func runLabel(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return errors.New("Only one bug id is supported")
	}

	if len(args) == 0 {
		return errors.New("You must provide a bug id")
	}

	backend, err := cache.NewRepoCache(repo)
	if err != nil {
		return err
	}
	defer backend.Close()

	prefix := args[0]

	b, err := backend.ResolveBugPrefix(prefix)
	if err != nil {
		return err
	}

	snap := b.Snapshot()

	for _, l := range snap.Labels {
		fmt.Println(l)
	}

	return nil
}

var labelCmd = &cobra.Command{
	Use:   "label <id>",
	Short: "Display a bug labels",
	RunE:  runLabel,
}

func init() {
	RootCmd.AddCommand(labelCmd)

	labelCmd.Flags().SortFlags = false
}
