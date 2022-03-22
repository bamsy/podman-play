/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Downs a pod using a kube.yaml file",
	Long: `Examples:
	- podman-play down (defaults to the kube.yaml file in current directory)
	- podman-play down /path/to/yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			downService(args[0])
		} else {
			downService(DEFAULT_KUBE_YAML)
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}

func downService(kubeYaml string) {
	fmt.Println(kubeYaml)

	// Start the service using the kube yaml file passed in
	cmd := exec.Command("podman", "play", "kube", kubeYaml, "--down")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
}
