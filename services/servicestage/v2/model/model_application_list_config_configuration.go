package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用配置信息。
type ApplicationListConfigConfiguration struct {
	// 应用组件环境变量。

	Env *[]ApplicationListConfigConfigurationEnv `json:"env,omitempty"`
}

func (o ApplicationListConfigConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationListConfigConfiguration struct{}"
	}

	return strings.Join([]string{"ApplicationListConfigConfiguration", string(data)}, " ")
}
