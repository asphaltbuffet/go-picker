package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/asphaltbuffet/go-picker/pkg/pick"

	log "github.com/sirupsen/logrus"
)

const (
	MinSource int = 2 // MinSource is the minimum number of items to pick from
)

var oneCmd *cobra.Command

// NewOneCommand returns the one command for the CLI.
func NewOneCommand() *cobra.Command {
	oneCmd = &cobra.Command{
		Use:  "one <args>...",
		Long: "one returns one random element from a list",
		Args: cobra.MinimumNArgs(MinSource),
		Run:  runOneCmd,
	}

	return oneCmd
}

func runOneCmd(cmd *cobra.Command, args []string) {
	log.WithField("src_len", len(args)).Trace("picking one element from input")

	p := pick.NewPicker()

	out, err := pick.GetOne(p, args)
	if err != nil {
		cmd.PrintErrln("failed to pick element: ", err)

		return
	}

	fmt.Println(out)
}
