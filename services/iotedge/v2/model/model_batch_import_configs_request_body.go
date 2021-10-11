package model

import (
	"encoding/json"

	"strings"
)

type BatchImportConfigsRequestBody struct {
	// 南向IA配置项列表

	Configs *[]BatchImportConfigRequestBody `json:"configs,omitempty"`
}

func (o BatchImportConfigsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchImportConfigsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchImportConfigsRequestBody", string(data)}, " ")
}
