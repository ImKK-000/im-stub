package config_test

import (
	. "im-stub/config"
	"im-stub/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadConfig_Should_Be_Error(t *testing.T) {
	_, actualError := ReadConfig("config.json")

	assert.Nil(t, actualError)
}
func Test_ReadConfig_Should_Be_Configs_Without_Error(t *testing.T) {
	expectedConfigs := model.Configs{
		Host:      "127.0.0.1",
		Port:      9999,
		EnableLog: false,
	}

	actualConfigs, actualError := ReadConfig("../../config/config.json")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedConfigs, actualConfigs)
}
