package dao

import (
	"techtrainingcamp-appUpgrade-team4/moudle"
	"techtrainingcamp-appUpgrade-team4/utils"
)

func QueryForOne(username string, password string) (*moudle.User, error) {

	sqlStr := "select id, username, password from t_user where username=? and password=?"

	row := utils.DB.QueryRow(sqlStr, username, password)
	user := &moudle.User{}
	row.Scan(&user.ID, &user.Username, &user.Password)
	return user, nil
}
