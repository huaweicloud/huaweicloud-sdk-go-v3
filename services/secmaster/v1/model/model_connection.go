package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Connection struct {

	// 数值
	ChannelReferCount *int64 `json:"channel_refer_count,omitempty"`

	// UUID
	ConnectionId *string `json:"connection_id,omitempty"`

	ConnectionType *FilterConnectionType `json:"connection_type,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 描述信息
	Info *string `json:"info,omitempty"`

	// UUID
	ModuleId *string `json:"module_id,omitempty"`

	// 相关标题
	TemplateTitle *string `json:"template_title,omitempty"`

	// 相关标题
	Title *string `json:"title,omitempty"`
}

func (o Connection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Connection struct{}"
	}

	return strings.Join([]string{"Connection", string(data)}, " ")
}
