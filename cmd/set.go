package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the status in Slack",
	Long:  "usage: slack-status set :emoji: [status]\n",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		e := args[0]

		var validEmojiSyntax = regexp.MustCompile(`:[\w\d].*:`)
		if validEmojiSyntax.MatchString(e) == false {
			e = fmt.Sprintf(":%s:", e)
		}

		s := strings.Join(args[1:], " ")

		c := slack.New(os.Getenv("SLACK_TOKEN"))
		err := c.SetUserCustomStatus(s, e)
		if err != nil {
			fmt.Printf("Failed: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
