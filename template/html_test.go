package template1

import (
	"html/template"
	"net/http"
	"os"
	"testing"
)

func tmpl(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("test.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")
}

func TestHtml(t *testing.T) {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/tmpl", tmpl)
	server.ListenAndServe()
}

func TestH2(t *testing.T)  {
	t1 := template.New("test1")
	tmpl, _ := t1.Parse(
		`
{{- define "T1"}}ONE {{println .}}{{end}}
{{- define "T2"}}{{template "T1" $}}{{end}}
{{- template "T2" . -}}
`)
	_ = tmpl.Execute(os.Stdout, "hello world")
}

func TestH3(t *testing.T)  {
	tx := template.Must(template.New("hh").Parse(
		`{{range $x := . -}}
{{println $x}}
{{- end}}
`))
	s := []int{11, 22, 33, 44, 55}
	_ = tx.Execute(os.Stdout, s)
}