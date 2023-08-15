package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationEnvParam struct {

	// 环境变量名称。
	Name string `json:"name"`

	// 环境变量取值。
	Value string `json:"value"`
}

func (o ConfigurationEnvParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationEnvParam struct{}"
	}

	return strings.Join([]string{"ConfigurationEnvParam", string(data)}, " ")
}
