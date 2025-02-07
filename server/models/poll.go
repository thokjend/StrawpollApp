package models

type Poll struct {
	Title       string   `json:"Title"`
	Description string   `json:"Description"`
	Options     []string `json:"Options"`
	Settings    []bool   `json:"Settings"`
}