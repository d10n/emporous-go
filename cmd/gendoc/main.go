package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	"github.com/emporous/emporous-go/cmd/client/commands"
)

func main() {
	genDocCmd := &cobra.Command{
		Use:          "gendoc",
		Short:        "Generate Emporous client CLI docs",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			cmd := commands.NewRootCmd()
			cmd.DisableAutoGenTag = true
			return doc.GenMarkdownTree(cmd, args[0])
		},
	}
	cobra.CheckErr(genDocCmd.Execute())
}
