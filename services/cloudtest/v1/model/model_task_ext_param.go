package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskExtParam struct {

	// 是否删除
	Delete *bool `json:"delete,omitempty"`

	// 参数id
	Id *string `json:"id,omitempty"`

	// 参数名称
	Name *string `json:"name,omitempty"`

	// 是否敏感信息：true-敏感信息，false-非敏感信息
	SensitiveInfo *bool `json:"sensitiveInfo,omitempty"`

	// 参数值
	Value *string `json:"value,omitempty"`

	// 参数类型
	VariableType *string `json:"variableType,omitempty"`
}

func (o TaskExtParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskExtParam struct{}"
	}

	return strings.Join([]string{"TaskExtParam", string(data)}, " ")
}
