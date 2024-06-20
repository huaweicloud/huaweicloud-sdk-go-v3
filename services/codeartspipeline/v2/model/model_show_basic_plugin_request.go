package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBasicPluginRequest Request Object
type ShowBasicPluginRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 插件名
	PluginName string `json:"plugin_name"`

	// 版本
	Version string `json:"version"`
}

func (o ShowBasicPluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBasicPluginRequest struct{}"
	}

	return strings.Join([]string{"ShowBasicPluginRequest", string(data)}, " ")
}
