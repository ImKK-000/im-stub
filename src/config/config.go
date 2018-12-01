package config

import (
	"encoding/json"
	"im-stub/model"
	"io/ioutil"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func validateData(data interface{}) error {
	validate = validator.New()
	return validate.Struct(data)
}

func ReadConfig(configPath string) (model.Configs, error) {
	rawConfigs, errorReadConfigFile := ioutil.ReadFile(configPath)
	if errorReadConfigFile != nil {
		return model.Configs{}, errorReadConfigFile
	}
	var configs model.Configs
	errorConvertJSONToStruct := json.Unmarshal(rawConfigs, &configs)
	if errorConvertJSONToStruct != nil {
		return model.Configs{}, errorConvertJSONToStruct
	}
	errorValidateData := validateData(configs)
	if errorConvertJSONToStruct != nil {
		return model.Configs{}, errorValidateData
	}
	return configs, nil
}
