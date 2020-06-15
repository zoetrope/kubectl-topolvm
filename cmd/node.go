package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/zoetrope/kubectl-topolvm/pkg"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// nodesCmd represents the nodes command
var nodesCmd = &cobra.Command{
	Use:   "node",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cli, err := pkg.KubernetesClient(cmd.PersistentFlags())
		if err != nil {
			return err
		}
		nodes := corev1.NodeList{}
		err = cli.List(context.Background(), &nodes, &client.ListOptions{})
		if err != nil {
			return err
		}
		summary, err := pkg.Summarize(cli)
		if err != nil {
			return err
		}
		err = pkg.PrintSummary(summary)
		return err
	},
}

func init() {
	rootCmd.AddCommand(nodesCmd)
}
