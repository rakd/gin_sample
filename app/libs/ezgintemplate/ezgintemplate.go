package ezgintemplate

import (
	"html/template"
	"log"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin/render"
)

// Render ...
type Render struct {
	Templates       map[string]*template.Template
	TemplatesDir    string
	Layout          string
	AmpLayout       string
	AdminLayout     string
	Ext             string
	TemplateFuncMap map[string]interface{}
	Debug           bool
	PartialsDir     string
}

// New ...
func New() Render {
	r := Render{

		// PartialsDir holds the location of the partials templates. which all tepmlates can use.
		PartialsDir: "partials/",

		Templates: map[string]*template.Template{},
		// TemplatesDir holds the location of the templates
		TemplatesDir: "app/views/",
		// Layout is the file name of the layout file
		Layout:      "layouts/base",
		AmpLayout:   "layouts/amp",
		AdminLayout: "layouts/admin",
		// Ext is the file extension of the rendered templates
		Ext: ".html",
		// Template's function map
		TemplateFuncMap: nil,
		// Debug enables debug mode
		Debug: false,
	}

	return r
}

// Init ...
func (r Render) Init() Render {
	layout := r.TemplatesDir + r.Layout + r.Ext
	ampLayout := r.TemplatesDir + r.AmpLayout + r.Ext
	adminLayout := r.TemplatesDir + r.AdminLayout + r.Ext

	viewDirs, _ := filepath.Glob(r.TemplatesDir + "**/*" + r.Ext)

	partials, err := filepath.Glob(r.TemplatesDir + r.PartialsDir + "*" + r.Ext)
	if err != nil {
		log.Print("cannot find " + r.PartialsDir)
		panic(err.Error())
	}

	for _, view := range viewDirs {
		renderName := r.getRenderName(view)
		if r.Debug {
			log.Printf("[GIN-debug] %-6s %-25s --> %s\n", "LOAD", view, renderName)
		}
		tpls := append([]string{layout, ampLayout, adminLayout, view}, partials...)
		r.AddFromFiles(renderName, tpls...)
		//r.AddFromFiles(renderName, layout, view)
	}

	return r
}

func (r Render) getRenderName(tpl string) string {
	dir, file := filepath.Split(tpl)
	dir = strings.Replace(dir, r.TemplatesDir, "", 1)
	file = strings.TrimSuffix(file, r.Ext)
	return dir + file
}

// Add ...
func (r Render) Add(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template can not be nil")
	}
	if len(name) == 0 {
		panic("template name cannot be empty")
	}
	r.Templates[name] = tmpl
}

// AddFromFiles ...
func (r Render) AddFromFiles(name string, files ...string) *template.Template {
	if strings.Contains(name, "_amp") {
		tmpl := template.Must(template.New(filepath.Base(r.AmpLayout + r.Ext)).Funcs(r.TemplateFuncMap).ParseFiles(files...))
		r.Add(name, tmpl)
		return tmpl
	} else if strings.Contains(name, "admin") {
		tmpl := template.Must(template.New(filepath.Base(r.AdminLayout + r.Ext)).Funcs(r.TemplateFuncMap).ParseFiles(files...))
		r.Add(name, tmpl)
		return tmpl
	}

	tmpl := template.Must(template.New(filepath.Base(r.Layout + r.Ext)).Funcs(r.TemplateFuncMap).ParseFiles(files...))
	r.Add(name, tmpl)
	return tmpl
}

// Instance ...
func (r Render) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r.Templates[name],
		Data:     data,
	}
}
