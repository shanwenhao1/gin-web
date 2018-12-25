package models

// gorm映射user表
type User struct {
	UserId   string      `json:"user_id" gorm:"primary_key"`
	UserName string      `json:"user_name" gorm:"column:user_name"`
	RegDate  interface{} `json:"reg_date" gorm:"column:reg_date"`
	Sex      int32       `json:"sex" gorm:"-"`
}
