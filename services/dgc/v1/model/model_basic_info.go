package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicInfo struct {

	// 作业责任人
	Owner *string `json:"owner,omitempty"`

	// 作业优先级
	Priority *int32 `json:"priority,omitempty"`

	// 作业执行用户
	ExecuteUser *string `json:"executeUser,omitempty"`

	// 实例超时时间
	InstanceTimeout *int32 `json:"instanceTimeout,omitempty"`

	// 用户自定义属性字段
	CustomFields *interface{} `json:"customFields,omitempty"`
}

func (o BasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicInfo struct{}"
	}

	return strings.Join([]string{"BasicInfo", string(data)}, " ")
}
