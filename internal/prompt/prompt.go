package prompt

import (
	"bytes"
	"text/template"
)

type Prompt struct {
	Code         string
	Version      string
	System       string
	UserTemplate string
}

func (p Prompt) Render(params map[string]string) (string, error) {
	tpl, err := template.New("user").Parse(p.UserTemplate)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, params)
	return buf.String(), err
}
