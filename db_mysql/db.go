package db_mysql

import "crypto/md5"

func AddUser(u models.User)(int64, error){
	md5Hash := md5.New()
	md5Hash.Write()
}