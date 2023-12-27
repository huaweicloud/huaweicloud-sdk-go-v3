package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BasicInfo 作业基本信息
type BasicInfo struct {

	// 作业责任人
	Owner *string `json:"owner,omitempty"`

	// 作业优先级，0代表高优先级，1代表中优先级，2代表低优先级。
	Priority *int32 `json:"priority,omitempty"`

	// 作业执行用户，必须是已存在的用户名。
	ExecuteUser *string `json:"execute_user,omitempty"`

	// 实例超时时间，单位是分钟。
	InstanceTimeout *int32 `json:"instance_timeout,omitempty"`

	// 用户自定义属性字段
	CustomFields *interface{} `json:"custom_fields,omitempty"`
}

func (o BasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicInfo struct{}"
	}

	return strings.Join([]string{"BasicInfo", string(data)}, " ")
}
