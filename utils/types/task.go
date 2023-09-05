package types

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Parameter struct {
	Name  string
	Value interface{} `gorm:"serializer:json"`
}

type Task struct {
	gorm.Model
	Targets    []string `gorm:"serializer:json"`
	TargetAll  bool
	Expire     time.Time
	Stop       time.Time
	ModuleID   string
	Parameters []Parameter `gorm:"serializer:json"`
	Log        []Log       `gorm:"serializer:json"`
	Output     []Output    `gorm:"serializer:json"`
}

type Output struct {
	Element string
	Data    interface{}
}

type Log struct {
	Time  time.Time
	Level uint8 // Debug, Info, Warn, Error, Fatal. (0-4)
	Body  string
}

func GetParameter(name string, parameters []Parameter) (Parameter, error) {
	for _, param := range parameters {
		if param.Name == name {
			return param, nil
		}
	}
	return Parameter{}, fmt.Errorf("could not find parameter: '%s'", name)
}
