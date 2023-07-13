package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceCreationVariables 变量值，通过模板创建时必填
type InstanceCreationVariables struct {

	// 环境变量
	Environment *interface{} `json:"environment,omitempty"`

	// 组件变量
	Components *interface{} `json:"components,omitempty"`
}

func (o InstanceCreationVariables) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceCreationVariables struct{}"
	}

	return strings.Join([]string{"InstanceCreationVariables", string(data)}, " ")
}
