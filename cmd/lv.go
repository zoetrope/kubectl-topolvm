package cmd

import (
	"context"

	topolvmv1 "github.com/cybozu-go/topolvm/api/v1"
	"github.com/spf13/cobra"
	"github.com/zoetrope/kubectl-topolvm/pkg"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// lvCmd represents the lv command
var lvCmd = &cobra.Command{
	Use:   "lv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cli, err := pkg.LogicalVolumeClient(cmd.PersistentFlags())
		if err != nil {
			return err
		}

		var lvlist topolvmv1.LogicalVolumeList
		err = cli.List(context.Background(), &lvlist, &client.ListOptions{})
		if err != nil {
			return err
		}

		return pkg.PrintLVList(&lvlist)
	},
}

func init() {
	rootCmd.AddCommand(lvCmd)
}
