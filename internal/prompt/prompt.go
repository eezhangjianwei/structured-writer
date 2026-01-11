// internal/prompt/prompt.go
package prompt

import (
	"bytes"
	"text/template"
)

type Prompt struct {
	System string
	Scene  string
	Format string
}

func (p Prompt) Render(params map[string]string) (string, error) {
	full := p.System + "\n\n" + p.Scene + "\n\n" + p.Format

	tpl, err := template.New("prompt").Parse(full)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, params)
	return buf.String(), err
}
