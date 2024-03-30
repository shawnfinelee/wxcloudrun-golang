package dao

import "wxcloudrun-golang/db/model"

type PlayerInterface interface {
	GetPlayer(id int32) (*model.PlayerModel, error)
}

type PlayerInterfaceImp struct{}

var PlayerImp PlayerInterface = &PlayerInterfaceImp{}
