package template1

import (
	"html/template"
	"os"
	"testing"

	"github.com/abc463774475/bbtool/n_log"
)

type Person1 struct {
	Name string
	Age  int
}

func TestF1(t *testing.T) {
	p := &Person1{"longshuai", 19}
	tmpl, err := template.New("test").Parse("Name:{{.Name}}, Age: {{.Age}}")
	n_log.Info("err  %v", err)

	err = tmpl.Execute(os.Stdout, p)
	n_log.Info("err  %v", err)

	n_log.Info("p %+v", p)
}

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func TestF2(t11 *testing.T) {
	f1 := Friend{Fname: "xiaofang"}
	f2 := Friend{Fname: "wugui"}
	t := template.New("test")
	t = template.Must(t.Parse(
		`hello {{.UserName}}!
{{ range .Emails }}
an email {{ . }}
{{- end }}
{{ with .Friends }}
{{- range . }}
my friend name is {{.Fname}}
{{- end }}
{{ end }}`))
	p := Person{
		UserName: "longshuai",
		Emails:   []string{"a1@qq.com", "a2@gmail.com"},
		Friends:  []*Friend{&f1, &f2},
	}
	t.Execute(os.Stdout, p)
}

