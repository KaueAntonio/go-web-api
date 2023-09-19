package models

type Pessoa struct {
	ID    int    `json: "id"`
	Name  string `json: "name"`
	Email string `json: "email"`
	Pass  string `json: "pass"`
}
