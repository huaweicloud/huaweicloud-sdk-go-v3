package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBasicPluginRequest Request Object
type DeleteBasicPluginRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 需要删除的插件名
	PluginName string `json:"plugin_name"`

	// 删除类型，all 代表删除整个插件，single代表删除单个插件版本
	Type string `json:"type"`

	// 需要删除的插件版本
	Version *string `json:"version,omitempty"`
}

func (o DeleteBasicPluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBasicPluginRequest struct{}"
	}

	return strings.Join([]string{"DeleteBasicPluginRequest", string(data)}, " ")
}
