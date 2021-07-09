package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMysqlJobInfoRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 任务ID。

	Id string `json:"id"`
}

func (o ShowMysqlJobInfoRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlJobInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowMysqlJobInfoRequest", string(data)}, " ")
}
