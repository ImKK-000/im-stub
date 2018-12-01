package model

type Configs struct {
	Host      string        `json:"host" validate:"required"`
	Port      uint          `json:"port" validate:"required"`
	EnableLog bool          `json:"enable_log,omitempty"`
	RouteList []ConfigRoute `json:"routes,omitempty"`
}

type ConfigRoute struct {
	Path     string      `json:"path" validate:"required"`
	Method   string      `json:"method" validate:"required"`
	Request  interface{} `json:"request" validate:"required"`
	Response interface{} `json:"response" validate:"required"`
}
