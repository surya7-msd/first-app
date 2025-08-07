package src

type User struct {
	Name string
	Id   string
	Age  int
}

func NewUser(u *User) *User {
	user := &User{
		Name: u.Name,
		Id:   u.Id,
		Age:  u.Age,
	}

	return user
}
