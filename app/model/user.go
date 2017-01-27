package model

type User struct {
	Id    int    `xorm:"not null pk autoincr INT(11)"`
	Email string `xorm:"not null unique VARCHAR(255)"`
	Name  string `xorm:"not null VARCHAR(255)"`
	Age   int    `xorm:"default 0 INT(11)"`
}
