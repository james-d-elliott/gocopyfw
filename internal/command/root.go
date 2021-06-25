package command

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/james-d-elliott/gocopyfw/internal/util"
)

// NewRootCmd returns the root cobra cmd.
func NewRootCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "gocopyfw",
		Short:   "Copy File Watcher for IntelliJ that copies files",
		Long:    rootCmdLong,
		Example: rootCmdExample,
		RunE:    runRootCmdE,
	}

	cmd.Flags().StringP("project-file-dir", "d", "", "the project root directory from $ProjectFileDir$")
	cmd.Flags().StringP("file-path", "f", "", "the file path from $FilePath$")
	cmd.Flags().StringSliceP("linked-files", "l", []string{}, "a list of files that are linked, each set of files is separated by commas and each file in the set is separated by a semicolon")

	_ = cmd.MarkFlagRequired("project-file-dir")
	_ = cmd.MarkFlagRequired("file-path")
	_ = cmd.MarkFlagRequired("linked-files")

	cmd.AddCommand(newCompletionCmd(), newMarkdownCmd())

	return cmd
}

func runRootCmdE(cmd *cobra.Command, _ []string) (err error) {
	pfd, _ := cmd.Flags().GetString("project-file-dir")
	fp, _ := cmd.Flags().GetString("file-path")

	fileName := strings.TrimPrefix(fp, pfd)

	fileSets, _ := cmd.Flags().GetStringSlice("linked-files")

	for _, fileSet := range fileSets {
		files := strings.Split(fileSet, ";")

		if util.IsStringInSlice(fileName, files) {
			src, err := os.Open(filepath.Join(pfd, fileName))
			if err != nil {
				return fmt.Errorf("could not open source file: %w", err)
			}

			for _, file := range files {
				if file != fileName {
					dest, err := os.OpenFile(filepath.Join(pfd, file), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0660)
					if err != nil {
						return fmt.Errorf("could not open destination file: %w", err)
					}

					_, err = io.Copy(dest, src)
					if err != nil {
						return fmt.Errorf("could not write destination file: %w", err)
					}
				}
			}
		}
	}

	return nil
}
