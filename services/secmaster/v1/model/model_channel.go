package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Channel struct {
	Action *ChannelAction `json:"action,omitempty"`

	// UUID
	ChannelId *string `json:"channel_id,omitempty"`

	ConfigStatus *ConfigStatus `json:"config_status,omitempty"`

	// Iam用户ID
	CreateBy *string `json:"create_by,omitempty"`

	// Iam用户名称
	CreateByUser *string `json:"create_by_user,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	ErrorType *ChannelErrorType `json:"error_type,omitempty"`

	HealthStatus *HealthStatus `json:"health_status,omitempty"`

	// UUID
	InputId *string `json:"input_id,omitempty"`

	// 所属租户名称
	InputName *string `json:"input_name,omitempty"`

	InstallStatus *InstallStatus `json:"install_status,omitempty"`

	// 毫秒时间戳
	InstallTime *int64 `json:"install_time,omitempty"`

	// 关联的节点个数
	NodeReferCount *int64 `json:"node_refer_count,omitempty"`

	OperationStatus *ChannelOperationStatus `json:"operation_status,omitempty"`

	// UUID
	OutputId *string `json:"output_id,omitempty"`

	// 所属租户名称
	OutputName *string `json:"output_name,omitempty"`

	// UUID
	ParserId *string `json:"parser_id,omitempty"`

	// 所属租户名称
	ParserName *string `json:"parser_name,omitempty"`

	ReadWrite *ReadWrite `json:"read_write,omitempty"`

	// 相关标题
	Title *string `json:"title,omitempty"`

	// 毫秒时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o Channel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Channel struct{}"
	}

	return strings.Join([]string{"Channel", string(data)}, " ")
}
