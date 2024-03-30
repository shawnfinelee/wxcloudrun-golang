package model

import (
	"time"
)

type PlayerModel struct {
	Id        int32     `gorm:"column:id" json:"id"`
	NickName  string    `gorm:"column:nickName" json:"nickName"`
	AvatarUrl string    `gorm:"column:avatarUrl" json:"avatarUrl"`
	Gender    int8      `gorm:"column:gender" json:"gender"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (PlayerModel) TableName() string {
	return "Player"
}
