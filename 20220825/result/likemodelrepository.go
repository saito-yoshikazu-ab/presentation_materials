package modelrepository

//go:generate mockgen -destinition=mock_$GOFILE -package=$GOPACKAGE

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Like struct {
	UserID    int64
	PostID    int64
	CreatedAt time.Time
}

type LikePK struct {
	UserID int64
	PostID int64
}

func (l *Like) toLikePK() LikePK {
	return LikePK{
		UserID: l.UserID,
		PostID: l.PostID,
	}
}

type Likes []*Like

func (ls *Likes) ToMap() map[LikePK]*Like {
	m := make(map[LikePK]*Like, len(*ls))
	for _, l := range *ls {
		m[l.toLikePK()] = l
	}
	return m
}

type LikeModelRepository struct {
	client *sqlx.DB
}

func (r *LikeModelRepository) Get(lk LikePK) (*Like, error) {
	model := new(Like)
	if err := r.client.Select(&model, `SELECT * FROM like WHERE
    UserID=?
    AND
    PostID=?
    `,
		lk.UserID,
		lk.PostID,
	); err != nil {
		return nil, err
	}
	return model, nil
}

func (r *LikeModelRepository) FindByUserID(userid string) (Likes, error) {
	var models Likes
	if err := r.client.Select(&models, "SELECT * FROM like WHERE userid=?", userid); err != nil {
		return nil, err
	}
	return models, nil
}

func (r *LikeModelRepository) FindByPostID(postid string) (Likes, error) {
	var models Likes
	if err := r.client.Select(&models, "SELECT * FROM like WHERE postid=?", postid); err != nil {
		return nil, err
	}
	return models, nil
}
