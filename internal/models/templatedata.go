package models

import "github.com/saroyangor/bookings/internal/forms"

type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	IntFloat        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	Form            *forms.Form
	IsAuthenticated int
}
