package model

import (
	"encoding/json"

	"strings"
)

type BasicInfo struct {
	// 作业责任人

	Owner *string `json:"owner,omitempty"`
	// 作业优先级

	Priority *string `json:"priority,omitempty"`
	// 作业执行用户

	ExecuteUser *string `json:"executeUser,omitempty"`
	// 实例超时时间

	InstanceTimeout *string `json:"instanceTimeout,omitempty"`
	// 用户自定义属性字段

	CustomFields *interface{} `json:"customFields,omitempty"`
}

func (o BasicInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BasicInfo struct{}"
	}

	return strings.Join([]string{"BasicInfo", string(data)}, " ")
}
