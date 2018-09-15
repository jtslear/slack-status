package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

func help() {
	fmt.Printf("usage: %s :emoji: [status]\n", os.Args[0])
	os.Exit(0)
}

func main() {
	a := os.Args

	if len(os.Args) <= 1 {
		help()
	}

	e := a[1]

	var validEmojiSyntax = regexp.MustCompile(`:[\w\d].*:`)
	if validEmojiSyntax.MatchString(e) == false {
		e = fmt.Sprintf(":%s:", e)
	}

	s := strings.Join(a[2:], " ")

	c := slack.New(os.Getenv("SLACK_TOKEN"))
	err := c.SetUserCustomStatus(s, e)
	if err != nil {
		fmt.Printf("Failed: %v\n", err)
	}
}
