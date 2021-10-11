package model

import (
	"encoding/json"

	"strings"
)

type ConfirmIaConfigRequestBody struct {
	// 配置项ID

	Id string `json:"id"`
	// 版本号

	Version string `json:"version"`
}

func (o ConfirmIaConfigRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ConfirmIaConfigRequestBody struct{}"
	}

	return strings.Join([]string{"ConfirmIaConfigRequestBody", string(data)}, " ")
}
