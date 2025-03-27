package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigBody 配置参数body
type ConfigBody struct {

	// 配置类型，分为：”EXCLUDE_MIGRATE_PATH\"，\"SYNC_EXCLUDE_PATH\"，\"ONLY_SYNC_PATH\"等
	ConfigKey string `json:"config_key"`

	// 具体配置参数字段，保存于数据库，最终在agent端进行解析
	ConfigValue string `json:"config_value"`

	// 描述配置状态的保留字段
	ConfigStatus *string `json:"config_status,omitempty"`
}

func (o ConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigBody struct{}"
	}

	return strings.Join([]string{"ConfigBody", string(data)}, " ")
}
