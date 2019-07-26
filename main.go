package main

// Response of flnow plugin spec for Golang
type Response struct {
	Result  bool        `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Raw     interface{} `json:"raw"`
}

// Field to descript field of a Form
type Field struct {
	Display     string `json:"display"`
	Name        string `json:"name"`
	ShortName   string `json:"short"`
	EnvName     string `json:"env"`
	Default     string `json:"default"`
	Description string `json:"description"`
	Value       string `json:"value"`
	Required    bool   `json:"required"`
	Type        string `json:"type"`
	Array       bool   `json:"array"`
}

// Form for flnow plugin UI display
type Form struct {
	Fields []Field `json:"fields"`
}
