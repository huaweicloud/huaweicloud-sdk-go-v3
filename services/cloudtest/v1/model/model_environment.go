package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Environment struct {

	// 环境分组id
	EnvironmentId *string `json:"environment_id,omitempty" xml:"environment_id"`

	// 环境分组名
	EnvironmentName *string `json:"environment_name,omitempty" xml:"environment_name"`

	// 环境分组描述
	EnvironmentDescription *string `json:"environment_description,omitempty" xml:"environment_description"`

	// 是否是默认环境
	IsDefault *bool `json:"is_default,omitempty" xml:"is_default"`
}

func (o Environment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Environment struct{}"
	}

	return strings.Join([]string{"Environment", string(data)}, " ")
}
