package user

import "errors"

type AuthServiceInterface interface {
	GetAllUser() (*[]User, error)
	GetSingleUser() (*User, error)
}

type UserService struct {
	repoService *repoStruct
}

func (userService *UserService) GetAllUser() (*[]User, error) {
	user, err := userService.repoService.GetAllUser()
	if err != nil {
		return &[]User{}, errors.New("couldn't retrieve data from database")
	}

	return user, nil
}

func (userService *UserService) GetSingleUser(username string) (*User, error) {
	user, err := userService.repoService.GetSingleUser(username)
	if err != nil {
		return &User{}, errors.New("couldn't retrieve data from database")
	}

	return user, nil
}

func NewUserService(repo *repoStruct) *UserService {
	return &UserService{
		repoService: repo,
	}
}
