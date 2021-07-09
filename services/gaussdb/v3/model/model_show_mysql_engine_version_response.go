package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMysqlEngineVersionResponse struct {
	// 数据库版本信息列表

	Datastores     *[]MysqlEngineVersionInfo `json:"datastores,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowMysqlEngineVersionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlEngineVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowMysqlEngineVersionResponse", string(data)}, " ")
}
