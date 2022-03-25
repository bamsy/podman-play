/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull all images in a kube.yaml file",
	Long: `Examples:
	- podman-play pull (defaults to the kube.yaml file in current directory)
	- podman-play pull /path/to/yaml.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			pullImages(args[0])
		} else {
			pullImages(DEFAULT_KUBE_YAML)
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func pullImages(kubeYaml string) {
	lines, _ := readLines(kubeYaml)
	const imageYamlKeyWord string = "image:"

	for _, s := range lines {
		words := strings.Split(strings.TrimLeft(s, " "), " ")
		if words[0] == imageYamlKeyWord {
			pullImage(words[1])
		}
	}
}

// Pull an image using the podman command
func pullImage(image string) {
	fmt.Printf("pulling image %s\n", image)
	cmd := exec.Command("podman", "pull", image)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(stdout))
		return
	}
	fmt.Println(string(stdout))
}

// Read our yaml file into a string array
func readLines(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}
