package user

type User interface {
	FetchName(id string) UserType
}

type UserType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
