package user

type IUserRepository interface {
	Create(u *User) error
	GetAll() (*[]User, error)
	GetOneById(id int) (*User, error)
	DeleteByID(id int) error
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

func (r *UserRepository) GetOneById(id int) (*User, error) {
	var user User

	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) DeleteByID(id int) error {
	if err := db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
