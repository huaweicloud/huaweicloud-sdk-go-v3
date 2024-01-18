package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicConfig struct {

	// 作业责任人
	Owner *string `json:"owner,omitempty"`

	// 作业委托的名称
	Agency *string `json:"agency,omitempty"`

	// 实例超时是否忽略等待时间, 取值范围为0和1, 0：表示实例超时不忽略等待时间1：表示实例超时忽略等待时间
	IsIgnoreWaiting *int32 `json:"isIgnoreWaiting,omitempty"`

	// 作业优先级
	Priority *int32 `json:"priority,omitempty"`

	// 作业执行用户
	ExecuteUser *string `json:"executeUser,omitempty"`

	// 实例超时时间
	InstanceTimeout *int32 `json:"instanceTimeout,omitempty"`

	// 用户自定义属性字段
	CustomFields *interface{} `json:"customFields,omitempty"`
}

func (o BasicConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicConfig struct{}"
	}

	return strings.Join([]string{"BasicConfig", string(data)}, " ")
}
