package domain

type Contact struct {
	Id        string `json:"Id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Status    string `json:"status"`
}
