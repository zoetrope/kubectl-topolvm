package cmd

import (
	"context"
	"fmt"

	topolvmv1 "github.com/cybozu-go/topolvm/api/v1"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
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
		flagSet := cmd.PersistentFlags()
		cfgFlags := genericclioptions.NewConfigFlags(true)
		cfgFlags.AddFlags(flagSet)
		matchVersionFlags := cmdutil.NewMatchVersionFlags(cfgFlags)
		factory := cmdutil.NewFactory(matchVersionFlags)

		restCfg, err := factory.ToRESTConfig()
		if err != nil {
			return err
		}
		crScheme := runtime.NewScheme()
		err = topolvmv1.AddToScheme(crScheme)
		if err != nil {
			return err
		}
		cli, err := client.New(restCfg, client.Options{Scheme: crScheme})
		if err != nil {
			return err
		}

		var lvlist topolvmv1.LogicalVolumeList
		err = cli.List(context.Background(), &lvlist, &client.ListOptions{})
		if err != nil {
			return err
		}

		for _, lv := range lvlist.Items {
			fmt.Println(lv.Name)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(lvCmd)
}
