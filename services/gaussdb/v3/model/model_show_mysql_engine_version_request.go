package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMysqlEngineVersionRequest struct {
	// 语言。

	XLanguage *string `json:"X-Language,omitempty"`
	// 数据库引擎名称。

	DatabaseName string `json:"database_name"`
}

func (o ShowMysqlEngineVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlEngineVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowMysqlEngineVersionRequest", string(data)}, " ")
}
