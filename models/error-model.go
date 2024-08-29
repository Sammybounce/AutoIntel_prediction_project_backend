package model

type CaptureError struct {
	Id        string `json:"id"`
	Error     string `json:"error"`
	CodeBlock string `json:"codeBlock"`
	FilePath  string `json:"filePath"`
	Env       string `json:"env"`
	CreatedAt string `json:"createdAt"`
}
