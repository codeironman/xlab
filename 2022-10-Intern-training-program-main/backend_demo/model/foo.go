package model
import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type Foo struct {
	Bar uint64 `gorm:"not null;autoIncrement;primaryKey"`
	Buz string `gorm:"not null"`
}

type struct User{
	id uint
	name string
	password string 
}
type struct Todos{
	id uint
	user_id uint
	title string
	content string
}
func ininMySql(){
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(todo), &gorm.Config{})
}
func main() {
	db.AutoMigrate(&Users{},&todos{})
	user:={11,"cfy","sss"}
	db.Create(user)//增加用户
	target_user:=db.First(&user)//查找用户
	db.Model(&user).Update("password","aaa")//更新用户信息
	db.Where(id=11).Delete(&User{})//删除用户


}