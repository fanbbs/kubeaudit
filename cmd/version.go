package cmd

import (
	"os"

	"github.com/hashicorp/go-version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Placeholder values will be overridden by goreleaser or makefile.
var (
	Version   = "0.0.0"
	Commit    = "ffffffff"
	BuildDate = "2006-01-02T15:04:05Z07:00"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kubeaudit",
	Long:  `This prints the version numbers of kubeaudit and the kubernetes server.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := version.NewVersion(Version + "+" + Commit)
		if err != nil {
			log.Fatal(err)
		}
		log.WithFields(log.Fields{
			"BuildDate": BuildDate,
			"Commit":    Commit,
			"Version":   Version,
		}).Info("Kubeaudit")

		kubeconfig := rootConfig.kubeConfig
		if rootConfig.localMode {
			kubeconfig = os.Getenv("HOME") + "/.kube/config"
		}

		kube, err := kubeClient(kubeconfig)
		if err != nil {
			log.Fatal(err)
		}
		printKubernetesVersion(kube)
	},
}
