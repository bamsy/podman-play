/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Starts a pod using a kube.yaml file",
	Long: `Examples:
	- podman-play up (defaults to the kube.yaml file in current directory)
	- podman-play up /path/to/yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			upService(args[0])
		} else {
			upService(DEFAULT_KUBE_YAML)
		}

	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}

func upService(kubeYaml string) {
	fmt.Println(kubeYaml)

	// Start the service using the kube yaml file passed in
	cmd := exec.Command("podman", "play", "kube", kubeYaml, "--log-driver", "journald")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
}
