package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	IntMap map[string]int
	Data   map[string]interface{}
}
