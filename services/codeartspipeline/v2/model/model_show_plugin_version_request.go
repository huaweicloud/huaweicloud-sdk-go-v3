package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginVersionRequest Request Object
type ShowPluginVersionRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 插件名
	PluginName string `json:"plugin_name"`

	// 版本
	Version *string `json:"version,omitempty"`
}

func (o ShowPluginVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowPluginVersionRequest", string(data)}, " ")
}
