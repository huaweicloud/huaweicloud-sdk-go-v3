package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 流程参数
type NextflowParamsDto struct {

	// 参数名
	Name string `json:"name"`

	// 参数值
	Value *string `json:"value,omitempty"`

	// 参数类型,取值[Other|File|Directory]
	Type *string `json:"type,omitempty"`

	// 参数描述。取值范围：[0-255]
	Description *string `json:"description,omitempty"`

	// 参数是否必填
	Required *bool `json:"required,omitempty"`
}

func (o NextflowParamsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NextflowParamsDto struct{}"
	}

	return strings.Join([]string{"NextflowParamsDto", string(data)}, " ")
}
