package main

import (
	"context"
	"database/sql"

	"github.com/superjcd/go-zero-templates/model/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// 需要自行将 dsn 中的 host，账号 密码配置正确
	dsn := "root:root@tcp(192.168.0.77:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local"
	conn := sqlx.NewMysql(dsn)
	_ = conn

	UserModel := user.NewUserModel(conn)
	result, _ := UserModel.Insert(context.Background(), &user.User{Name: sql.NullString{String: "testuser"}, Password: "123456"})
	_ = result

}
