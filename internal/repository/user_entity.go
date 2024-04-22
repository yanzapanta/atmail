package repository

type User struct {
	ID       uint
	Username string
	Email    string
	Age      int
}

func (User) TableName() string {
	return "users"
}
