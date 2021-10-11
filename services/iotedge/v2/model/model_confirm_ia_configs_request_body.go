package model

import (
	"encoding/json"

	"strings"
)

type ConfirmIaConfigsRequestBody struct {
	// 确认配置项列表

	Configs *[]ConfirmIaConfigRequestBody `json:"configs,omitempty"`
}

func (o ConfirmIaConfigsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ConfirmIaConfigsRequestBody struct{}"
	}

	return strings.Join([]string{"ConfirmIaConfigsRequestBody", string(data)}, " ")
}
