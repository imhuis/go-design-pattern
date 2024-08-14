package example

type IUser interface {
	Login(username, password string) (*User, error)
	Register(username, password string) (*User, error)
}

// IUserFacade facade interface,actually expose to outer
type IUserFacade interface {
	LoginOrRegister(username, password string) (*User, error)
}

type User struct {
	Name   string
	Passwd string
}

type UserService struct{}

func (u *UserService) LoginOrRegister(username, password string) (*User, error) {
	user, err := u.Login(username, password)
	if err != nil {
		return u.Register(username, password)
	}
	return user, nil
}

func (u *UserService) Login(username, password string) (*User, error) {
	return &User{Name: ""}, nil
}

func (u *UserService) Register(username, password string) (*User, error) {
	return &User{Name: ""}, nil
}
