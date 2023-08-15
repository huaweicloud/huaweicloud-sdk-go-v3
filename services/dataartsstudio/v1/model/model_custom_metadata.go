package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomMetadata struct {

	// 描述
	Description *string `json:"description,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 任务id
	TaskId *string `json:"task_id,omitempty"`
}

func (o CustomMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomMetadata struct{}"
	}

	return strings.Join([]string{"CustomMetadata", string(data)}, " ")
}
