package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentParam struct {

	// 组件描述
	Description *string `json:"description,omitempty"`

	// 应用Id、子应用Id
	ModelId string `json:"model_id"`

	// 应用、子应用，取值：APPLICATION、SUB_APPLICATION
	ModelType string `json:"model_type"`

	// 组件名称
	Name string `json:"name"`
}

func (o ComponentParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentParam struct{}"
	}

	return strings.Join([]string{"ComponentParam", string(data)}, " ")
}
