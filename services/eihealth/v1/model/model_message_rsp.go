package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 消息邮件发送配置
type MessageRsp struct {

	// 消息类型
	MessageType *string `json:"message_type,omitempty"`

	// 项目名称
	EihealthProjectName *string `json:"eihealth_project_name,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 操作用户
	Operator *string `json:"operator,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 操作时间
	OperateTime *string `json:"operate_time,omitempty"`

	// 详情
	MessageDetail *string `json:"message_detail,omitempty"`
}

func (o MessageRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MessageRsp struct{}"
	}

	return strings.Join([]string{"MessageRsp", string(data)}, " ")
}
