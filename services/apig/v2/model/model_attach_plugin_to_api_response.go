package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachPluginToApiResponse Response Object
type AttachPluginToApiResponse struct {

	// 绑定插件信息列表。
	AttachedPlugins *[]PluginApiAttachInfo `json:"attached_plugins,omitempty"`
	HttpStatusCode  int                    `json:"-"`
}

func (o AttachPluginToApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachPluginToApiResponse struct{}"
	}

	return strings.Join([]string{"AttachPluginToApiResponse", string(data)}, " ")
}
