package model

type Configs struct {
	Host      string `json:"host" validate:"required,ip"`
	Port      uint   `json:"port" validate:"required"`
	EnableLog bool   `json:"enable_log,omitempty"`
}
