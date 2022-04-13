package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用配置信息。
type ApplicationConfigModifyConfiguration struct {
	// 应用环境变量。

	Env []ApplicationConfigModifyConfigurationEnv `json:"env"`
}

func (o ApplicationConfigModifyConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationConfigModifyConfiguration struct{}"
	}

	return strings.Join([]string{"ApplicationConfigModifyConfiguration", string(data)}, " ")
}
