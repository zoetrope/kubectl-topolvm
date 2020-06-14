package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zoetrope/kubectl-topolvm/pkg"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
		client, err := pkg.KubernetesClient(cmd.PersistentFlags())
		if err != nil {
			return err
		}
		nodes, err := client.CoreV1().Nodes().List(metav1.ListOptions{})
		if err != nil {
			return err
		}
		for _, node := range nodes.Items {
			fmt.Println(node.Name)
			for k, v := range node.Annotations {
				if strings.HasPrefix(k, "capacity.topolvm.cybozu.com/") {
					dc := k[len("capacity.topolvm.cybozu.com/"):]
					bytes, err := strconv.ParseUint(v, 10, 64)
					if err != nil {
						return err
					}
					fmt.Printf("  %s: %s\n", dc, pkg.FormatBytes(bytes))
				}
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(nodesCmd)
}
