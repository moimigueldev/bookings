package models

// Template data holds data sent from the handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // interface means any data can be put there
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
