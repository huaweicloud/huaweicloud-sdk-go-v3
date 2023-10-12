package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachApiToPluginResponse Response Object
type AttachApiToPluginResponse struct {

	// 绑定插件信息列表。
	AttachedPlugins *[]PluginApiAttachInfo `json:"attached_plugins,omitempty"`
	HttpStatusCode  int                    `json:"-"`
}

func (o AttachApiToPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachApiToPluginResponse struct{}"
	}

	return strings.Join([]string{"AttachApiToPluginResponse", string(data)}, " ")
}
