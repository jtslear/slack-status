# slack-status
* Sets one's Slack Status via the command line

## Requirements
This is utilizing the Slack oauth Apps Token.  One must create an application in
slack and install the app into the desired workspace.  When creating the
application it'll need the `users.profile:write` permission scope.

Take the token provided by slack and set this in to the `SLACK_TOKEN`
environment variable.

## Installation
* `go install github.com/jtslear/slack-status`

## Usage
* `slack-status :lunch: out to eat`

## Development
* This is using the [go11 modules experiment](https://github.com/golang/go/wiki/Modules#go-111-modules)
* So you'll want go version 1.11+
* fork this repo
* clone the fork
* `cd` into wherever you cloned it
* `go build`
* have fun
