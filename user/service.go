package user

type AuthServiceInterface interface {
	GetAllUser() (*[]User, error)
	GetSingleUser() (*User, error)
	GetAllUserWithTask() (*UserResponse, error)
	GetSingleUserWithTask() (*Data, error)
}

type UserService struct {
	repoService *repoStruct
}

func (userService *UserService) GetAllUser() (*[]Data, error) {
	user, err := userService.repoService.getAllUser()
	if err != nil {
		return &[]Data{}, err
	}

	return user, nil
}

func (userService *UserService) GetSingleUser(id string) (*User, error) {
	user, err := userService.repoService.getSingleUser(id)
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (userService *UserService) GetAllUserWithTask() (*[]Data, error) {
	user, err := userService.repoService.getAllUserWithTask()
	if err != nil {
		return &[]Data{}, err
	}

	return user, nil
}

func (userService *UserService) GetSingleUserWithTask(id string) (*Data, error) {
	user, err := userService.repoService.getSingleUserWithTask(id)
	if err != nil {
		return &Data{}, err
	}

	return user, nil
}

func NewUserService(repo *repoStruct) *UserService {
	return &UserService{
		repoService: repo,
	}
}
