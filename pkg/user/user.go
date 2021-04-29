package user

type User interface {
	Fetch(id string) UserType
}

type UserType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
