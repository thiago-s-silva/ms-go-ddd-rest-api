package user

type IUserRepository interface {
	Create(u *User) error
	GetAll() (*[]User, error)
}

type UserRepository struct{}

func (r *UserRepository) Create(u *User) error {
	res := db.Create(&u)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *UserRepository) GetAll() (*[]User, error) {
	var users []User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}
