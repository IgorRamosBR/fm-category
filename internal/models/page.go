package models

type Page struct {
	Results []Category `json:"results"`
	Next    string     `json:"next"`
}
