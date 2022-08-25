package modelrepository

//go:generate mockgen -destinition=mock_$GOFILE -package=$GOPACKAGE

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Post struct {
	ID        int64
	UserID    int64
	Content   string
	CreatedAt time.Time
}

type PostPK struct {
	ID int64
}

func (p *Post) toPostPK() PostPK {
	return PostPK{
		ID: p.ID,
	}
}

type Posts []*Post

func (ps *Posts) ToMap() map[PostPK]*Post {
	m := make(map[PostPK]*Post, len(*ps))
	for _, p := range *ps {
		m[p.toPostPK()] = p
	}
	return m
}

type PostModelRepository struct {
	client *sqlx.DB
}

func (r *PostModelRepository) Get(pk PostPK) (*Post, error) {
	model := new(Post)
	if err := r.client.Select(&model, `SELECT * FROM post WHERE
    ID=?
    `,
		pk.ID,
	); err != nil {
		return nil, err
	}
	return model, nil
}

func (r *PostModelRepository) FindByUserID(userid string) (Posts, error) {
	var models Posts
	if err := r.client.Select(&models, "SELECT * FROM post WHERE userid=?", userid); err != nil {
		return nil, err
	}
	return models, nil
}
