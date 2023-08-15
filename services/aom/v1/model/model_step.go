package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Step 创建作业时的步骤参数
type Step struct {

	// 步骤id。
	Id *string `json:"id,omitempty"`

	// 步骤名称。
	Name *string `json:"name,omitempty"`

	// 步骤类型。
	Type *string `json:"type,omitempty"`

	// 步骤参数。
	Input map[string]string `json:"input,omitempty"`

	// 是否自动忽略错误。
	IgnoreError bool `json:"ignore_error"`

	// 步骤说明。
	Description *string `json:"description,omitempty"`
}

func (o Step) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Step struct{}"
	}

	return strings.Join([]string{"Step", string(data)}, " ")
}
