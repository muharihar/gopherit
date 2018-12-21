package ui

import (
	"00-newapp-template/internal/pkg"
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
)

type CLI struct {
	Config *pkg.Config
}

func NewCLI(c *pkg.Config) (cli CLI) {
	cli.Config = c
	return
}

func (cli *CLI) DrawGopher() {
	fmt.Println(Gopher())
	return
}

func Gopher() string {
	gopher := `
	         ,_---~~~~~----._         
	  _,,_,*^____      _____''*g*\"*, 
	 / __/ /'     ^.  /      \ ^@q   f 
	[  @f | @))    |  | @))   l  0 _/  
	 \'/   \~____ / __ \_____/    \   
	  |           _l__l_           I   
	  }          [______]           I  
	  ]            | | |            |  
	  ]             ~ ~             |  
	  |                            |   
	
	[[@https://gist.github.com/belbomemo]]
	`
	return gopher
}

func (cli *CLI) Render(name string, data interface{}) (usage string) {
	var raw bytes.Buffer
	var err error

	templateDir:= cli.Config.TemplateFolder
	templateDir = strings.TrimSuffix(templateDir,"/")

	t := template.New("")
	t, err = t.Funcs(
		template.FuncMap{
			"Gopher": Gopher,
		},
	).ParseGlob(fmt.Sprintf("%s/ui/*.tmpl", templateDir))

	if err != nil {
		log.Fatalf("couldn't load template: %v", err)
	}

	err = t.ExecuteTemplate(&raw, name, data)
	if err != nil {
		log.Fatalf("error in Execute template: %v", err)
	}

	usage = raw.String()
	return
}
