package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tableNamePlayer = "Player"

// GetPlayer implements PlayerInterface.
func (*PlayerInterfaceImp) GetPlayer(id int32) (*model.PlayerModel, error) {
	var err error
	var player = new(model.PlayerModel)

	cli := db.Get()
	err = cli.Table(tableNamePlayer).Where("id = ?", id).First(player).Error

	return player, err
}
