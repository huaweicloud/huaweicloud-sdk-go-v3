package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyRedisConfigBody 修改配置参数请求体
type ModifyRedisConfigBody struct {

	// 实例配置项数组。
	RedisConfig *[]RedisConfig `json:"redis_config,omitempty"`
}

func (o ModifyRedisConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyRedisConfigBody struct{}"
	}

	return strings.Join([]string{"ModifyRedisConfigBody", string(data)}, " ")
}
