package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailableConfig struct {

	// 自定义关键字是否开通
	CustomAwAvailable *bool `json:"custom_aw_available,omitempty"`

	// 系统关键字是否开通
	PublicAwAvailable *bool `json:"public_aw_available,omitempty"`

	// 一键刷新功能是否开通
	RefreshAwAvailable *bool `json:"refresh_aw_available,omitempty"`
}

func (o AvailableConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableConfig struct{}"
	}

	return strings.Join([]string{"AvailableConfig", string(data)}, " ")
}
