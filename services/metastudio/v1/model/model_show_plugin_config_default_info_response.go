package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginConfigDefaultInfoResponse Response Object
type ShowPluginConfigDefaultInfoResponse struct {

	// 默认插件支持列表
	DefaultPluginList *[]DefaultPluginInfo `json:"default_plugin_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPluginConfigDefaultInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginConfigDefaultInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowPluginConfigDefaultInfoResponse", string(data)}, " ")
}
