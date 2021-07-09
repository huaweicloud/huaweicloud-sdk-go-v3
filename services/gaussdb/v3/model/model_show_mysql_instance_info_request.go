package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMysqlInstanceInfoRequest struct {
	// 语言。

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowMysqlInstanceInfoRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlInstanceInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowMysqlInstanceInfoRequest", string(data)}, " ")
}
