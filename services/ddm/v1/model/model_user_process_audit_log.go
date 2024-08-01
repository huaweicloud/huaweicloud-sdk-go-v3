package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserProcessAuditLog struct {

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名
	InstanceName *string `json:"instance_name,omitempty"`

	// 会话id
	ProcessId *string `json:"process_id,omitempty"`

	// 执行用户名
	ExecuteUserName *string `json:"execute_user_name,omitempty"`

	// 发生时间，UTC时间
	ExcuteTime *string `json:"excute_time,omitempty"`
}

func (o UserProcessAuditLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserProcessAuditLog struct{}"
	}

	return strings.Join([]string{"UserProcessAuditLog", string(data)}, " ")
}
