package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 场景元数据。
type SceneAssetMeta struct {

	// 可操作组件列表（如屏幕，灯光，摄像机）。
	Components *[]SceneComponentInfo `json:"components,omitempty"`

	// 默认场景设置（机位，初始人位置）。
	DefaultConfigs map[string]SceneComponentInfo `json:"default_configs,omitempty"`
}

func (o SceneAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SceneAssetMeta struct{}"
	}

	return strings.Join([]string{"SceneAssetMeta", string(data)}, " ")
}
