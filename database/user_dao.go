package database

import (
	"errors"

	"acgfate/model"
)

type UserDao struct{}

// QueryByUID query single row by UID
func (d *UserDao) QueryByUID(idx interface{}) (*model.User, error) {
	var user model.User
	sqlStr := "SELECT * FROM af_user WHERE uid = ?"
	err := db.Get(&user, sqlStr, idx)
	return &user, err
}

// QueryByUname query single row by username
func (d *UserDao) QueryByUname(idx interface{}) (*model.User, error) {
	var user model.User
	sqlStr := "SELECT * FROM af_user WHERE username = ?"
	err := db.Get(&user, sqlStr, idx)
	return &user, err
}

// QueryByEmail query single row by email
func (d *UserDao) QueryByEmail(idx interface{}) (*model.User, error) {
	var user model.User
	sqlStr := "SELECT * FROM af_user WHERE email = ?"
	err := db.Get(&user, sqlStr, idx)
	return &user, err
}

func (d *UserDao) InsertRow(userRow interface{}) error {
	if user, ok := userRow.(model.User); ok {
		tx, err := db.Begin()
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		sqlStr := "INSERT INTO af_user(username, password, nickname, email) VALUES (?, ?, ?, ?)"
		ret, err := tx.Exec(sqlStr, user.Username, user.Password, user.Nickname, user.Email)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		// get last insert row id
		uid, err := ret.LastInsertId()
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		// insert into info table
		infoDao := UserInfoDao{}
		if err := infoDao.InsertRow(tx, uid); err != nil {
			_ = tx.Rollback()
			return err
		}
		_ = tx.Commit()
		return nil
	}
	return errors.New("user model incorrect")
}

func (d *UserDao) UpdateRow() {}

func (d *UserDao) DeleteRow() {}

func (d *UserDao) DeleteMRow() {}
