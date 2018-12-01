package config

import (
	"im-stub/model"
)

func ReadConfig() (model.Configs, error) {
	return model.Configs{
		Host: "127.0.0.1",
		Port: 9999,
		Log:  false,
	}, nil
}
