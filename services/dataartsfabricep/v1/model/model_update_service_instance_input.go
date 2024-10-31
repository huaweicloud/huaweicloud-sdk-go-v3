package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServiceInstanceInput 更新Service的请求体
type UpdateServiceInstanceInput struct {
	Source *SourceRef `json:"source,omitempty"`

	// 一个Service Instance的名称，只能包含中文、字母、数字、下划线、中划线、点、空格
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	Config *UpdateServiceInstanceConfig `json:"config,omitempty"`
}

func (o UpdateServiceInstanceInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceInstanceInput struct{}"
	}

	return strings.Join([]string{"UpdateServiceInstanceInput", string(data)}, " ")
}
