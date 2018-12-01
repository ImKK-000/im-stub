package config_test

import (
	"fmt"
	. "im-stub/config"
	"im-stub/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadConfig_Should_Be_Read_File_Error(t *testing.T) {
	filename := "xxx"
	expectedErrorMessage := fmt.Sprintf("open %s: no such file or directory", filename)

	_, actualError := ReadConfig(filename)

	assert.Equal(t, expectedErrorMessage, actualError.Error())
}

func Test_ReadConfig_Should_Be_Configs_Without_Error(t *testing.T) {
	expectedConfigs := model.Configs{
		Host:      "127.0.0.1",
		Port:      9999,
		EnableLog: false,
	}

	actualConfigs, actualError := ReadConfig("../../config/config.json")

	fmt.Println(actualError)
	assert.Nil(t, actualError)
	assert.Equal(t, expectedConfigs, actualConfigs)
}
