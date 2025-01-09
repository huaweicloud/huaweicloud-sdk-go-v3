package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CorpConfigInfo 企业配置。
type CorpConfigInfo struct {

	// 查询的配置configKey。
	ConfigKey *string `json:"config_key,omitempty"`

	// 根据configKey查到的configValue。
	ConfigValue *string `json:"config_value,omitempty"`
}

func (o CorpConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CorpConfigInfo struct{}"
	}

	return strings.Join([]string{"CorpConfigInfo", string(data)}, " ")
}
