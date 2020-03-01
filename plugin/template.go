package plugin

import (
	"bytes"
	"text/template"
)

// Default template
const DefaultTemplate = `{{ .BuildInfo.Status.Icon }} {{ .RepoInfo.ParsedName }}.
Build {{ .BuildInfo.Number }}  {{ .BuildInfo.Status.Message }}
{{ .CommitInfo.Author }} pushed {{ .CommitInfo.Sha }} to {{ .CommitInfo.Branch }} — {{ .CommitInfo.Message }}
{{ .BuildInfo.Link }}
`

// drone-telegram like template
const DroneTelegramTemplate = `{{ .BuildInfo.Status.Icon }} Build {{ .BuildInfo.Number }} of {{ .RepoInfo.ParsedName }} {{ .BuildInfo.Status.Message }}.
📝 Commit by {{ .CommitInfo.Author }} on {{ .CommitInfo.Branch }}:
	{{ .CommitInfo.Message }}

🌐 {{ .BuildInfo.Link }}
`

func ExecuteTemplate(tmpl string, info Info) (string, error) {
	t := template.New("template")

	t, err := t.Parse(tmpl)
	if err != nil {
		return "", err
	}

	b := new(bytes.Buffer)
	err = t.Execute(b, info)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
