package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginInfoListResponse Response Object
type ListPluginInfoListResponse struct {

	// 插件名称
	PluginName *string `json:"plugin_name,omitempty"`

	// 端口
	Port *string `json:"port,omitempty"`

	// 插件版本
	PluginVersion *string `json:"plugin_version,omitempty"`

	// 是否已安装
	Installed      *string `json:"installed,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPluginInfoListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginInfoListResponse struct{}"
	}

	return strings.Join([]string{"ListPluginInfoListResponse", string(data)}, " ")
}
