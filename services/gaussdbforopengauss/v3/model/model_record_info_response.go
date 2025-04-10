package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecordInfoResponse 操作记录详情
type RecordInfoResponse struct {

	// 主键id
	Id *string `json:"id,omitempty"`

	// 动作
	Action *string `json:"action,omitempty"`

	// 操作状态
	Status *string `json:"status,omitempty"`

	// 信息
	Message *string `json:"message,omitempty"`

	// 实体id
	EntityId *string `json:"entity_id,omitempty"`

	// 实体类型
	EntityType *string `json:"entity_type,omitempty"`

	// 工作流id
	JobId *string `json:"job_id,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 扩展信息
	ExtendedInfo *interface{} `json:"extended_info,omitempty"`
}

func (o RecordInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordInfoResponse struct{}"
	}

	return strings.Join([]string{"RecordInfoResponse", string(data)}, " ")
}
