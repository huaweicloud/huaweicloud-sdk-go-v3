package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicDrugModel struct {

	// 模型ID
	Id *string `json:"id,omitempty"`

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 模型名称
	Name *string `json:"name,omitempty"`

	// 模型创建者
	Creator *string `json:"creator,omitempty"`

	// 模型类型
	Type *string `json:"type,omitempty"`

	ValueRange *ValueRange `json:"value_range,omitempty"`

	// 模型描述信息
	Description *string `json:"description,omitempty"`
}

func (o BasicDrugModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicDrugModel struct{}"
	}

	return strings.Join([]string{"BasicDrugModel", string(data)}, " ")
}
