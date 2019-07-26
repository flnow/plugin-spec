package spec

// Response of flnow plugin spec for Golang
type Response struct {
	Result  bool        `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Raw     interface{} `json:"raw"`
}
