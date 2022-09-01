package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点重装场景服务器配置
type ReinstallServerConfig struct {

	// 云服务器标签，键必须唯一，CCE支持的最大用户自定义标签数量依region而定，自定义标签数上限为5个。
	UserTags *[]UserTag `json:"userTags,omitempty" xml:"userTags"`

	RootVolume *ReinstallVolumeSpec `json:"rootVolume,omitempty" xml:"rootVolume"`
}

func (o ReinstallServerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerConfig struct{}"
	}

	return strings.Join([]string{"ReinstallServerConfig", string(data)}, " ")
}
