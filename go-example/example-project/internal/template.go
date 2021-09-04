package internal

import (
	"log"
	"os"
	"text/template"
	"time"
)

const templ = `
		Number: {{.Number}}
		Title: {{.Title | printf "%.64s"}}
		Age: {{.CreatedAt | daysAgo}} days
		`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

type issue struct {
	Number    int32
	Title     string
	CreatedAt time.Time
}

func TemplateOperation() {

	result := issue{
		Number:    12,
		Title:     "Hello",
		CreatedAt: time.Now(),
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}
