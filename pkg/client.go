package pkg

import (
	topolvmv1 "github.com/cybozu-go/topolvm/api/v1"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func LogicalVolumeClient(flagSet *pflag.FlagSet) (client.Client, error) {
	cfgFlags := genericclioptions.NewConfigFlags(true)
	cfgFlags.AddFlags(flagSet)
	matchVersionFlags := cmdutil.NewMatchVersionFlags(cfgFlags)
	factory := cmdutil.NewFactory(matchVersionFlags)

	restCfg, err := factory.ToRESTConfig()
	if err != nil {
		return nil, err
	}
	crScheme := runtime.NewScheme()
	err = topolvmv1.AddToScheme(crScheme)
	if err != nil {
		return nil, err
	}
	cli, err := client.New(restCfg, client.Options{Scheme: crScheme})
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func KubernetesClient(flagSet *pflag.FlagSet) (*kubernetes.Clientset, error) {
	cfgFlags := genericclioptions.NewConfigFlags(true)
	cfgFlags.AddFlags(flagSet)
	matchVersionFlags := cmdutil.NewMatchVersionFlags(cfgFlags)
	factory := cmdutil.NewFactory(matchVersionFlags)
	cli, err := factory.KubernetesClientSet()
	if err != nil {
		return nil, err
	}
	return cli, err
}
