package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationConfigModifyConfigurationEnv struct {
	// 环境变量名称。

	Name string `json:"name"`
	// 环境变量取值。

	Value string `json:"value"`
}

func (o ApplicationConfigModifyConfigurationEnv) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationConfigModifyConfigurationEnv struct{}"
	}

	return strings.Join([]string{"ApplicationConfigModifyConfigurationEnv", string(data)}, " ")
}
