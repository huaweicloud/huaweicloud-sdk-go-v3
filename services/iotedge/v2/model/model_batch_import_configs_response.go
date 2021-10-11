package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchImportConfigsResponse struct {
	// 已成功导入的配置项id

	Ids            *interface{} `json:"ids,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchImportConfigsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchImportConfigsResponse struct{}"
	}

	return strings.Join([]string{"BatchImportConfigsResponse", string(data)}, " ")
}
