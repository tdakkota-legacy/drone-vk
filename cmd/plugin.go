package main

import (
	"drone-vk/plugin"
	"github.com/urfave/cli"
	"log"
	"os"
)

func pluginFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Usage:  "VK token",
			EnvVar: "VK_TOKEN,TOKEN,PLUGIN_TOKEN",
		},
		cli.IntFlag{
			Name:   "peer_id",
			Usage:  "User/Chat to send message",
			EnvVar: "VK_PEER_ID,PLUGIN_TO,PLUGIN_PEER_ID",
		},
		cli.StringFlag{
			Name:   "template",
			Usage:  "message text template",
			EnvVar: "VK_TEMPLATE,TEMPLATE,PLUGIN_TEMPLATE",
		},
		cli.StringFlag{
			Name:   "image",
			Usage:  "image to send",
			EnvVar: "VK_IMAGE,IMAGE,PLUGIN_IMAGE",
		},
		cli.StringFlag{
			Name:   "file",
			Usage:  "file to send",
			EnvVar: "VK_FILE,FILE,PLUGIN_FILE",
		},
		cli.StringFlag{
			Name:   "file.type",
			Usage:  "can be doc(default), graffiti, audio_message",
			EnvVar: "VK_FILE_TYPE,FILE_TYPE,PLUGIN_FILE_TYPE",
		},
		cli.IntFlag{
			Name:   "sticker",
			Usage:  "sticker(id) to send",
			EnvVar: "VK_STICKER,STICKER,PLUGIN_STICKER",
		},
		cli.StringFlag{
			Name:   "poll",
			Usage:  "poll to send",
			EnvVar: "VK_POLL,POLL,PLUGIN_POLL",
		},
	}
}

func droneFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "repo",
			Usage:  "repository owner and repository name",
			EnvVar: "DRONE_REPO,GITHUB_REPOSITORY",
		},
		cli.StringFlag{
			Name:   "repo.namespace",
			Usage:  "repository namespace",
			EnvVar: "DRONE_REPO_OWNER,DRONE_REPO_NAMESPACE,GITHUB_ACTOR",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA,GITHUB_SHA",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF,GITHUB_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.link",
			Usage:  "git commit link",
			EnvVar: "DRONE_COMMIT_LINK",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.author.email",
			Usage:  "git author email",
			EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
		},
		cli.StringFlag{
			Name:   "commit.author.avatar",
			Usage:  "git author avatar",
			EnvVar: "DRONE_COMMIT_AUTHOR_AVATAR",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.StringFlag{
			Name:   "pull.request",
			Usage:  "pull request",
			EnvVar: "DRONE_PULL_REQUEST",
		},
		cli.Float64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Float64Flag{
			Name:   "job.finished",
			Usage:  "job finished",
			EnvVar: "DRONE_BUILD_FINISHED",
		},
		cli.StringFlag{
			Name:   "deploy.to",
			Usage:  "Provides the target deployment environment for the running build. This value is only available to promotion and rollback pipelines.",
			EnvVar: "DRONE_DEPLOY_TO",
		},
	}
}

func githubFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:   "github",
			Usage:  "Boolean value, indicates the runtime environment is GitHub Action.",
			EnvVar: "PLUGIN_GITHUB,GITHUB",
		},
		cli.StringFlag{
			Name:   "github.workflow",
			Usage:  "The name of the workflow.",
			EnvVar: "GITHUB_WORKFLOW",
		},
		cli.StringFlag{
			Name:   "github.action",
			Usage:  "The name of the action.",
			EnvVar: "GITHUB_ACTION",
		},
		cli.StringFlag{
			Name:   "github.event.name",
			Usage:  "The webhook name of the event that triggered the workflow.",
			EnvVar: "GITHUB_EVENT_NAME",
		},
		cli.StringFlag{
			Name:   "github.event.path",
			Usage:  "The path to a file that contains the payload of the event that triggered the workflow. Value: /github/workflow/event.json.",
			EnvVar: "GITHUB_EVENT_PATH",
		},
		cli.StringFlag{
			Name:   "github.workspace",
			Usage:  "The GitHub workspace path. Value: /github/workspace.",
			EnvVar: "GITHUB_WORKSPACE",
		},
	}
}

func allFlags() []cli.Flag {
	flags := pluginFlags()
	flags = append(flags, droneFlags()...)
	flags = append(flags, githubFlags()...)
	return flags
}

func createApp(callback cli.ActionFunc) *cli.App {
	app := cli.NewApp()
	app.Name = "vk_plugin"
	app.Action = callback
	app.Flags = allFlags()
	return app
}

func main() {
	if err := createApp(plugin.Plugin{}.App).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
