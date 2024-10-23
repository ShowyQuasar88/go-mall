package main

import (

	"github.com/joho/godotenv"
	"github.com/showyquasar88/go-mall/demo/demo_proto/biz/dal"
	"github.com/showyquasar88/go-mall/demo/demo_proto/biz/dal/mysql"
	"github.com/showyquasar88/go-mall/demo/demo_proto/biz/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	// CURD

	// create
	// mysql.DB.Create(&model.User{
	// 	Email: "ginwithouta@gmail.com",
	// 	Password: "1234",
	// })

	// update
	// mysql.DB.Model(&model.User{}).Where("email = ?", "ginwithouta@gmail.com").Update("password", "1111")
	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email = ?", "ginwithouta@gmail.com").First(&row)
	// fmt.Println(row)

	// soft delete
	// mysql.DB.Where("email = ?", "ginwithouta@gmail.com").Delete(&model.User{})

	// hard delete
	mysql.DB.Unscoped().Where("email = ?", "ginwithouta@gmail.com").Delete(&model.User{})
}
