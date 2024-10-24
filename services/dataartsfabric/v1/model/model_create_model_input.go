package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelInput 创建模型的输入请求
type CreateModelInput struct {

	// 一个Model的名称，只能包含中文、字母、数字、下划线、中划线、点、空格
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	Type *ModelType `json:"type"`

	Version *ModelVersionInput `json:"version"`
}

func (o CreateModelInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelInput struct{}"
	}

	return strings.Join([]string{"CreateModelInput", string(data)}, " ")
}
