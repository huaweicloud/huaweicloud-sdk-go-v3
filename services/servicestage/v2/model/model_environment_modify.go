package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnvironmentModify struct {
	// 环境名称。

	Name *string `json:"name,omitempty"`
	// 环境别名。

	Alias *string `json:"alias,omitempty"`
	// 环境描述。

	Description *string `json:"description,omitempty"`
}

func (o EnvironmentModify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentModify struct{}"
	}

	return strings.Join([]string{"EnvironmentModify", string(data)}, " ")
}
