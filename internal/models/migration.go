package models

import (
	"errors"
)

type Migration struct{}

// 首次启用, 创建数据库表
func (migration *Migration) Install(dbName string) error {
	setting := new(Setting)
	task := new(Task)
	tables := []interface{}{
		&User{}, task, &TaskLog{}, &Host{}, setting, &LoginLog{}, &TaskHost{},
	}
	for _, table := range tables {
		exist, err := Db.IsTableExist(table)
		if exist {
			return errors.New("数据表已存在")
		}
		if err != nil {
			return err
		}
		err = Db.Sync2(table)
		if err != nil {
			return err
		}
	}
	setting.InitBasicField()

	return nil
}
