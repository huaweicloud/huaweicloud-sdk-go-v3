package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChQueryActionInfo 实例操作动作。
type ChQueryActionInfo struct {

	// 实例动作ID。
	Id *string `json:"id,omitempty"`

	// 实例动作名称。
	Action *string `json:"action,omitempty"`

	// 实例动作对象ID。
	ObjectId *string `json:"object_id,omitempty"`

	// 实例动作类型。
	Type *string `json:"type,omitempty"`

	// 实例动作任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 实例动作状态。
	Status *string `json:"status,omitempty"`

	// 实例动作创建时间。
	CreatedAt *int64 `json:"created_at,omitempty"`

	// 实例动作更新时间。
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o ChQueryActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChQueryActionInfo struct{}"
	}

	return strings.Join([]string{"ChQueryActionInfo", string(data)}, " ")
}
