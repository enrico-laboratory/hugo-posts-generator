package hugopostgenerator

import (
	"os"
	"path/filepath"
	"text/template"
	"time"
)

const templatesPath = "templates"

type templateData struct {
	Date          time.Time
	Description   string
	FeaturedImage string
	Tags          []string
	Title         string
	Year          string
	Details       []Details
	Content       string
}

type Details struct {
	Date   time.Time
	Venue  Venue
	Ticket Ticket
}

type Venue struct {
	Name    string
	Address string
}

type Ticket struct {
	Description string
	Link        string
	Price       string
}

func createTemplate(td *templateData) {
	tmpl, err := template.New("post").ParseFS(fs, filepath.Join(templatesPath, "post.tmpl"))
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "post", td)
	if err != nil {
		panic(err)
	}
}
