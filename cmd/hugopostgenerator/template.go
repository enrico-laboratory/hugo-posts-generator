package hugopostgenerator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

const templatesPath = "templates"

type TemplateData struct {
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

func (c *Client) GeneratePost(td *templateData, postName string) error {
	tmpl, err := createTemplate(td)
	if err != nil {
		return err
	}
	f, err := createFile(c.postFolder, postName)
	if err != nil {
		return err
	}
	err = tmpl.ExecuteTemplate(f, "post", td)
	if err != nil {
		return err
	}
	return nil
}

func createTemplate(td *templateData) (*template.Template, error) {
	tmpl, err := template.New("post").ParseFS(fs, filepath.Join(templatesPath, "post.tmpl"))
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func createFile(postFolder, postName string) (*os.File, error) {
	fileName := fmt.Sprintf("%s.md", postName)
	filePath := filepath.Join(postFolder, fileName)
	f, err := os.Create(filePath)
	return f, err
}
