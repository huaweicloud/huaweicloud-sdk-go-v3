package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Environment struct {

	// 环境分组id
	EnvironmentId *string `json:"environment_id,omitempty"`

	// 环境分组名
	EnvironmentName *string `json:"environment_name,omitempty"`

	// 环境分组描述
	EnvironmentDescription *string `json:"environment_description,omitempty"`

	// 是否是默认环境
	IsDefault *bool `json:"is_default,omitempty"`
}

func (o Environment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Environment struct{}"
	}

	return strings.Join([]string{"Environment", string(data)}, " ")
}
