package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RedisConfig struct {
	// 实例配置项的值。

	ParamValue string `json:"param_value"`
	// 实例配置项名。

	ParamName string `json:"param_name"`
	// 实例配置项ID。

	ParamId string `json:"param_id"`
}

func (o RedisConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisConfig struct{}"
	}

	return strings.Join([]string{"RedisConfig", string(data)}, " ")
}
