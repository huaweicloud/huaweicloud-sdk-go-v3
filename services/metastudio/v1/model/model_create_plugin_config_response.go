package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePluginConfigResponse Response Object
type CreatePluginConfigResponse struct {

	// 插件配置ID。
	PluginConfigId *string `json:"plugin_config_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePluginConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePluginConfigResponse struct{}"
	}

	return strings.Join([]string{"CreatePluginConfigResponse", string(data)}, " ")
}
