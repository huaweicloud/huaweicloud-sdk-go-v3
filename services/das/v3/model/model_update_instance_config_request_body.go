package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceConfigRequestBody 设置实例配置请求体
type UpdateInstanceConfigRequestBody struct {

	// 配置类型。取值范围：metaLockWaitThreshold, innodbLockWaitThreshold
	ConfigType string `json:"config_type"`

	// 配置的数值
	ConfigValue string `json:"config_value"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`
}

func (o UpdateInstanceConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigRequestBody", string(data)}, " ")
}
