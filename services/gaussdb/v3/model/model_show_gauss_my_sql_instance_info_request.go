package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowGaussMySqlInstanceInfoRequest struct {
	// 语言。

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowGaussMySqlInstanceInfoRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlInstanceInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlInstanceInfoRequest", string(data)}, " ")
}
