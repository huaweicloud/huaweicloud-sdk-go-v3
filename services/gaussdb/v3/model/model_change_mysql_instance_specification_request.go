package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ChangeMysqlInstanceSpecificationRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`

	Body *MysqlChangeSpecificationRequest `json:"body,omitempty"`
}

func (o ChangeMysqlInstanceSpecificationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeMysqlInstanceSpecificationRequest struct{}"
	}

	return strings.Join([]string{"ChangeMysqlInstanceSpecificationRequest", string(data)}, " ")
}
