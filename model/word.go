package model

import "time"

// Word 文字模型
type Word struct {
	ID        int64     `db:"id"`
	Wid       int64     `db:"wid"`
	Aid       int64     `db:"aid"`
	CatID     int64     `db:"cat_id"`
	Status    int64     `db:"status"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	Likes     int64     `db:"likes"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

const (
	StatusWordNormal  = 0
	StatusWordDeleted = 2
)

// IsDeleted return if word have been deleted.
func (w *Word) IsDeleted() bool {
	return w.Status == StatusWordDeleted
}

func (w *Word) UpdateLikes(val int64) {
	w.Likes = val
}
