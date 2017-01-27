package model

type UserItem struct {
	Id     int    `xorm:"not null pk autoincr INT(11)"`
	UserId int    `xorm:"not null index INT(11)"`
	Name   string `xorm:"default '' VARCHAR(64)"`
}
