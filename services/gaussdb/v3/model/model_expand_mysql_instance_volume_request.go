package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExpandMysqlInstanceVolumeRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`

	Body *MysqlExtendInstanceVolumeRequest `json:"body,omitempty"`
}

func (o ExpandMysqlInstanceVolumeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExpandMysqlInstanceVolumeRequest struct{}"
	}

	return strings.Join([]string{"ExpandMysqlInstanceVolumeRequest", string(data)}, " ")
}
