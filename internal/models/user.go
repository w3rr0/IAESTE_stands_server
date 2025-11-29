package models

type Name struct {
	First string
	Last  string
}

type Group string

const (
	IT      Group = "IT"
	PR      Group = "PR"
	IO      Group = "IO"
	JFR     Group = "JFR"
	Grafika Group = "Grafika"
)

type User struct {
	ID    int8
	Name  Name
	Email string
	Group Group
}
