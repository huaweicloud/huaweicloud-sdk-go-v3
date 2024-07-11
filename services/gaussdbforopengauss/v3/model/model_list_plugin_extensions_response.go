package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginExtensionsResponse Response Object
type ListPluginExtensionsResponse struct {

	// 实例插件拓展信息
	Body           *[]PluginExtensions `json:"body,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListPluginExtensionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginExtensionsResponse struct{}"
	}

	return strings.Join([]string{"ListPluginExtensionsResponse", string(data)}, " ")
}
