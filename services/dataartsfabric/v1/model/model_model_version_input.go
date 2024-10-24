package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelVersionInput 每个版本的详细信息
type ModelVersionInput struct {

	// 模型版本名称, 只能包含中文、字母、数字、下划线、中划线、点、空格
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	Config *ModelConfig `json:"config,omitempty"`
}

func (o ModelVersionInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelVersionInput struct{}"
	}

	return strings.Join([]string{"ModelVersionInput", string(data)}, " ")
}
