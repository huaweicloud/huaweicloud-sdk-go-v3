package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelInput 更新模型的输入请求
type UpdateModelInput struct {

	// 一个Model的名称，只能包含中文、字母、数字、下划线、中划线、点、空格
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 模型版本ID，32~36位的英文、数字、短横组合，系统自动生成无法修改，输入不生效。
	CurrentVersionId *string `json:"current_version_id,omitempty"`

	Version *ModelVersionInput `json:"version,omitempty"`
}

func (o UpdateModelInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelInput struct{}"
	}

	return strings.Join([]string{"UpdateModelInput", string(data)}, " ")
}
