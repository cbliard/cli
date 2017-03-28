package cmd

import (
	"github.com/Scalingo/cli/appdetect"
	"github.com/Scalingo/cli/cmd/autocomplete"
	"github.com/Scalingo/cli/deployments"
	"github.com/Scalingo/codegangsta-cli"
	"github.com/Scalingo/go-scalingo/io"
)

var (
	DeploymentsListCommand = cli.Command{
		Name:     "deployments",
		Category: "Deployment",
		Usage:    "List app deployments",
		Flags:    []cli.Flag{appFlag},
		Description: ` List all of your previous app deployments
    $ scalingo -a myapp deployments
`,
		Before: AuthenticateHook,
		Action: func(c *cli.Context) {
			currentApp := appdetect.CurrentApp(c)
			err := deployments.List(currentApp)
			if err != nil {
				errorQuit(err)
			}
		},
	}
	DeploymentLogCommand = cli.Command{
		Name:     "deployment-logs",
		Category: "Deployment",
		Usage:    "View deployment logs",
		Flags:    []cli.Flag{appFlag},
		Description: ` Get the logs of an app deployment
		$ scalingo -a myapp deployment-logs my-deployment
`,
		Before: AuthenticateHook,
		Action: func(c *cli.Context) {
			currentApp := appdetect.CurrentApp(c)
			if len(c.Args()) == 1 {
				err := deployments.Logs(currentApp, c.Args()[0])
				if err != nil {
					errorQuit(err)
				}
			} else {
				cli.ShowCommandHelp(c, "deployment-logs")
			}
		},
		BashComplete: func(c *cli.Context) {
			autocomplete.DeploymentsAutoComplete(c)
		},
	}
	DeploymentFollowCommand = cli.Command{
		Name:     "deployment-follow",
		Category: "Deployment",
		Usage:    "Follow deployment event stream",
		Flags:    []cli.Flag{appFlag},
		Description: ` Get real-time deployment informations
		$ scalingo -a myapp deployment-follow
`,
		Before: AuthenticateHook,
		Action: func(c *cli.Context) {
			currentApp := appdetect.CurrentApp(c)
			err := deployments.Stream(&deployments.StreamOpts{
				AppName: currentApp,
			})
			if err != nil {
				errorQuit(err)
			}
		},
	}
	DeploymentDeployCommand = cli.Command{
		Name:     "deploy",
		Category: "Deployment",
		Usage:    "Trigger a deployment by archive",
		Flags: []cli.Flag{appFlag,
			cli.BoolFlag{Name: "war, w", Usage: "Specify that you want to deploy a WAR file"},
		},
		Description: ` Trigger the deployment of a custom archive for your application
		$ scalingo -a myapp deploy archive.tar.gz
		or
		$ scalingo -a myapp deploy http://example.com/archive.tar.gz

    # See also commands 'deployments'
`,
		Before: AuthenticateHook,
		Action: func(c *cli.Context) {
			args := c.Args()
			if len(args) != 1 && len(args) != 2 {
				cli.ShowCommandHelp(c, "deploy")
				return
			}
			archivePath := args[0]
			gitRef := ""
			if len(args) == 2 {
				gitRef = args[1]
			}
			currentApp := appdetect.CurrentApp(c)
			if c.Bool("war") {
				io.Status("Deploy a WAR archive")
				err := deployments.DeployWar(currentApp, archivePath, gitRef)
				if err != nil {
					errorQuit(err)
				}
			} else {
				io.Status("Deploy an archive")
				err := deployments.Deploy(currentApp, archivePath, gitRef)
				if err != nil {
					errorQuit(err)
				}
			}
		},
		BashComplete: func(c *cli.Context) {
			autocomplete.CmdFlagsAutoComplete(c, "deploy")
		},
	}
)
