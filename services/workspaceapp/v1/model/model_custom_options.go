package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomOptions 自定义策略设置项。
type CustomOptions struct {

	// 配置项1内容。
	CustomConfiguration1Rule *string `json:"custom_configuration1_rule,omitempty"`
}

func (o CustomOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomOptions struct{}"
	}

	return strings.Join([]string{"CustomOptions", string(data)}, " ")
}
