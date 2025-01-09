package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesCustomOptions 通用音视频设置项。
type PoliciesCustomOptions struct {

	// 自定义配置规则。
	CustomConfiguration1Rule *string `json:"custom_configuration1_rule,omitempty"`
}

func (o PoliciesCustomOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesCustomOptions struct{}"
	}

	return strings.Join([]string{"PoliciesCustomOptions", string(data)}, " ")
}
