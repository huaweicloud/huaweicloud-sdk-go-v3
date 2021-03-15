package model

import (
	"encoding/json"

	"strings"
)

type SslOptionRequest struct {
	// - true, 打开ssl开关。 - false, 关闭ssl开关。
	SslOption bool `json:"ssl_option"`
}

func (o SslOptionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SslOptionRequest struct{}"
	}

	return strings.Join([]string{"SslOptionRequest", string(data)}, " ")
}
