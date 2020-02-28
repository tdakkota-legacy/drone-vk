package plugin

import (
	"github.com/urfave/cli"
	"strings"
)

// GitHubInfo struct
// Github information.
type GitHubInfo struct {
	Workflow  string
	Workspace string
	Action    string
	EventName string
	EventPath string
}

func (i *GitHubInfo) Parse(c *cli.Context) {
	i.Workflow = c.String("github.workflow")
	i.Workspace = c.String("github.workspace")
	i.Action = c.String("github.action")
	i.EventName = c.String("github.event.name")
	i.EventPath = c.String("github.event.path")
}

// RepoInfo struct
// Repo information.
type RepoInfo struct {
	FullName   string
	Namespace  string
	Name       string
	ParsedName string
}

func (i *RepoInfo) Parse(c *cli.Context) {
	i.FullName = c.String("repo")
	i.Namespace = c.String("repo.namespace")
	i.Name = c.String("repo.name")

	if i.FullName == "" {
		i.ParsedName = i.Namespace + "/" + i.Name
	} else {
		i.ParsedName = i.FullName
	}
}

// CommitInfo struct
// Commit information.
type CommitInfo struct {
	Sha     string
	Ref     string
	Branch  string
	Link    string
	Author  string
	Avatar  string
	Email   string
	Message string
}

func (i *CommitInfo) Parse(c *cli.Context) {
	i.Sha = c.String("commit.sha")
	i.Ref = c.String("commit.ref")
	i.Branch = c.String("commit.branch")
	i.Link = c.String("commit.link")
	i.Author = c.String("commit.author")
	i.Email = c.String("commit.author.email")
	i.Avatar = c.String("commit.author.avatar")
	i.Message = c.String("commit.message")
}

type Status struct {
	Name    string
	Icon    string
	Message string
}

func StatusFromString(s string) Status {
	switch strings.ToLower(s) {
	case "success":
		return Status{"Success", "✅", "succeeded"}
	case "failure":
		return Status{"Failure", "❌", "failed"}
	case "cancelled":
		return Status{"Cancelled", "❕", "cancelled"}
	default:
		return Status{}
	}
}

// BuildInfo struct
// Build information.
type BuildInfo struct {
	Tag      string
	Event    string
	Number   int
	Status   Status
	Link     string
	Started  float64
	Finished float64
	PR       string
	DeployTo string
}

func (i *BuildInfo) Parse(c *cli.Context) {
	i.Tag = c.String("build.tag")
	i.Number = c.Int("build.number")
	i.Event = c.String("build.event")
	i.Status = StatusFromString(c.String("build.status"))
	i.Link = c.String("build.link")
	i.Started = c.Float64("job.started")
	i.Finished = c.Float64("job.finished")
	i.PR = c.String("pull.request")
	i.DeployTo = c.String("deploy.to")
}

type Info struct {
	GitHubInfo GitHubInfo
	RepoInfo   RepoInfo
	CommitInfo CommitInfo
	BuildInfo  BuildInfo
}

func (i *Info) Parse(c *cli.Context) {
	i.GitHubInfo.Parse(c)
	i.RepoInfo.Parse(c)
	i.CommitInfo.Parse(c)
	i.BuildInfo.Parse(c)
}

func ParseInfo(c *cli.Context) (info Info) {
	info.Parse(c)
	return info
}
