package user

type AuthServiceInterface interface {
	GetAllUser() (*[]Data, error)
	GetAllUserWithTask() (*[]Data, error)
	GetSingleUser(id string) (*Data, error)
	GetSingleUserWithTask(id string) (*Data, error)
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

func (userService *UserService) GetAllUserWithTask() (*[]Data, error) {
	user, err := userService.repoService.getAllUserWithTask()
	if err != nil {
		return &[]Data{}, err
	}

	return user, nil
}

func (userService *UserService) GetSingleUser(id string) (*Data, error) {
	user, err := userService.repoService.getSingleUser(id)
	if err != nil {
		return &Data{}, err
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
