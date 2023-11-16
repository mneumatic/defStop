package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	Production    bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger

	//Session       *scs.SessionManager
	//Testimonials  interface{}
	//Products      interface{}
}
