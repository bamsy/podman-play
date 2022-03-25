/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// default path to kube yaml file
const DEFAULT_KUBE_YAML = "kube.yaml"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "podman-play",
	Short: "A wrapper around the 'podman play kube' to work similar to docker-compose",
	Long: `If a file is not passed in "kube.yaml" is used by default. Commands supported:
	- podman-play up -> Will start your pod using a passed in kube.yaml. 
	- podman-play down -> Will stop and your service using a passed in kube.yaml.
	- podman-play pull -> Will pull down any images in a passed in kube.yaml file.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.podman-play.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
