package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventSpecResponse 事件配置对象
type EventSpecResponse struct {

	// 事件配置ID
	Id *string `json:"id,omitempty"`

	// 事件配置定义名称
	Name *string `json:"name,omitempty"`

	// 事件配置显示名称
	DisplayName *string `json:"display_name,omitempty"`

	// 事件配置描述
	Description *string `json:"description,omitempty"`

	// 事件主题
	Subject *string `json:"subject,omitempty"`

	// 事件类别
	Category *string `json:"category,omitempty"`

	// 事件级别
	Severity *string `json:"severity,omitempty"`

	// 事件源类型
	SourceType *string `json:"source_type,omitempty"`

	// 所属服务
	NameSpace *string `json:"name_space,omitempty"`
}

func (o EventSpecResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventSpecResponse struct{}"
	}

	return strings.Join([]string{"EventSpecResponse", string(data)}, " ")
}
