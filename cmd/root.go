package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var urlRoot = "https://www.youtube.com/results?search_query="

var rootCmd = &cobra.Command{
	Use:   "youtube-search",
	Short: "Search quickly for a youtube video",
	Long:  "Youtube search is a way to quickly search for a youtube video in the command line.",
	Run: func(cmd *cobra.Command, args []string) {

		var query string

		// Build the search query here from the args.
		for _, arg := range args {
			query += fmt.Sprintf("%s ", arg)
		}

		searchUrl := urlRoot + query

		err := openBrowser(searchUrl)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// This should cover most OS
func openBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}
