package model

type UserInfo struct {
	ID     int64  `gorm:"column:id;primary_key" db:"id" json:"id"`
	Name   string `gorm:"column:name" db:"name" json:"name"`
	Gender int32  `gorm:"column:gender" db:"gender" json:"gender"`
	Age    int32  `gorm:"column:age" db:"age" json:"age"`
	Job    string `gorm:"column:job" db:"job" json:"job"`
}

func (userInfo *UserInfo) TableName() string {
	return "user_info"
}
