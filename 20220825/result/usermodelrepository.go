package modelrepository

//go:generate mockgen -destinition=mock_$GOFILE -package=$GOPACKAGE

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID            int64
	Name          string
	LastLoginDate time.Time
}

type UserPK struct {
	ID int64
}

func (u *User) toUserPK() UserPK {
	return UserPK{
		ID: u.ID,
	}
}

type Users []*User

func (us *Users) ToMap() map[UserPK]*User {
	m := make(map[UserPK]*User, len(*us))
	for _, u := range *us {
		m[u.toUserPK()] = u
	}
	return m
}

type UserModelRepository struct {
	client *sqlx.DB
}

func (r *UserModelRepository) Get(uk UserPK) (*User, error) {
	model := new(User)
	if err := r.client.Select(&model, `SELECT * FROM user WHERE
    ID=?
    `,
		uk.ID,
	); err != nil {
		return nil, err
	}
	return model, nil
}

func (r *UserModelRepository) FindByName(name string) (Users, error) {
	var models Users
	if err := r.client.Select(&models, "SELECT * FROM user WHERE name=?", name); err != nil {
		return nil, err
	}
	return models, nil
}
