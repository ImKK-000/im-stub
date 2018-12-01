package config

import (
	"encoding/json"
	"im-stub/model"
	"io/ioutil"

	"gopkg.in/go-playground/validator.v9"
)

var defaultConfigs = model.Configs{
	Host:      "0.0.0.0",
	Port:      9999,
	EnableLog: false,
}

func validateData(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}

func ReadConfig(configPath string) (model.Configs, error) {
	rawConfigs, errorReadConfigFile := ioutil.ReadFile(configPath)
	if errorReadConfigFile != nil {
		return defaultConfigs, errorReadConfigFile
	}
	var configs model.Configs
	errorConvertJSONToStruct := json.Unmarshal(rawConfigs, &configs)
	if errorConvertJSONToStruct != nil {
		return defaultConfigs, errorConvertJSONToStruct
	}
	errorValidateData := validateData(configs)
	if errorValidateData != nil {
		return defaultConfigs, errorValidateData
	}
	return configs, nil
}
