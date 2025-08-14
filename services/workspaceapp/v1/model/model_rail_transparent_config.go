package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RailTransparentConfig 应用镜像策略设置项。
type RailTransparentConfig struct {

	// 策略选值原则。
	SelectPolicy *int32 `json:"select_policy,omitempty"`

	// 配置项内容。
	TransparentListCustom *string `json:"transparent_list_custom,omitempty"`
}

func (o RailTransparentConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RailTransparentConfig struct{}"
	}

	return strings.Join([]string{"RailTransparentConfig", string(data)}, " ")
}
