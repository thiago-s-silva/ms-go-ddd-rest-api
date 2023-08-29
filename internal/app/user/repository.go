package user

type IUserRepository interface {
	Create(u *User) error
}

type UserRepository struct{}

func (r *UserRepository) Create(u *User) error {
	res := db.Create(&u)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
