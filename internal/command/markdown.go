package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func newMarkdownCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "markdown",
		Short: "Generate Markdown docs for the command",
		RunE:  runMarkdownCmdE,
	}

	cmd.Flags().StringP("directory", "d", "", "the directory for the docs")

	return cmd
}

func runMarkdownCmdE(cmd *cobra.Command, _ []string) (err error) {
	dir, err := cmd.Flags().GetString("directory")
	if err != nil {
		return err
	}

	return doc.GenMarkdownTree(cmd.Root(), dir)
}
