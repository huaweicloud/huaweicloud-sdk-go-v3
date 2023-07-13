package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePluginsReq struct {

	// 是否开启该插件。
	Enable *bool `json:"enable,omitempty"`

	// 插件列表，多个插件中间用“,”隔开。
	Plugins *string `json:"plugins,omitempty"`
}

func (o UpdatePluginsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePluginsReq struct{}"
	}

	return strings.Join([]string{"UpdatePluginsReq", string(data)}, " ")
}
