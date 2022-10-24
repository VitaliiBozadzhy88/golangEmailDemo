package service

type UserService interface {
	Insert(Email string, Code int)
	GetEmail(Email string) string
}
