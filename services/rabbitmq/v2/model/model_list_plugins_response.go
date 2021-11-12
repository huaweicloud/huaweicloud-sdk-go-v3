package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPluginsResponse struct {
	// 插件信息列表。

	Plugins        *[]PluginEntity `json:"plugins,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListPluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginsResponse struct{}"
	}

	return strings.Join([]string{"ListPluginsResponse", string(data)}, " ")
}
