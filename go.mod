module github.com/zoetrope/kubectl-topolvm

go 1.13

require (
	github.com/cybozu-go/topolvm v0.4.8
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	k8s.io/api v0.17.6
	k8s.io/apimachinery v0.17.6
	k8s.io/cli-runtime v0.17.6
	k8s.io/client-go v0.17.6
	k8s.io/kubectl v0.17.6
	sigs.k8s.io/controller-runtime v0.5.5
)
