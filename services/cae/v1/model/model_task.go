package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Task struct {

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 任务详情。
	Detail *string `json:"detail,omitempty"`

	// 任务序号。
	Index *int32 `json:"index,omitempty"`

	// 任务名称
	Name *string `json:"name,omitempty"`

	// 任务状态
	Status *string `json:"status,omitempty"`

	// 修改时间
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o Task) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Task struct{}"
	}

	return strings.Join([]string{"Task", string(data)}, " ")
}
