package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMessageRequest struct {

	// 医疗智能体项目名
	EihealthProjectName *string `json:"eihealth_project_name,omitempty"`

	// 查询条数
	Limit *int32 `json:"limit,omitempty"`

	// 消息类型
	MessageType *string `json:"message_type,omitempty"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 操作者名称
	Operator *string `json:"operator,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 操作状态
	Status *string `json:"status,omitempty"`
}

func (o ListMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageRequest struct{}"
	}

	return strings.Join([]string{"ListMessageRequest", string(data)}, " ")
}
