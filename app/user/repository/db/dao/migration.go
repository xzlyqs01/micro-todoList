package dao

import "micro-todoList/micro-todoList/app/user/repository/db/model"

func migration() {
	_db.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(&model.User{})
}
