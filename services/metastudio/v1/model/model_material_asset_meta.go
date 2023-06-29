package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MaterialAssetMeta 素材元数据。
type MaterialAssetMeta struct {

	// 可替换的素材组件列表。
	Components *[]MaterialComponentInfo `json:"components,omitempty"`
}

func (o MaterialAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaterialAssetMeta struct{}"
	}

	return strings.Join([]string{"MaterialAssetMeta", string(data)}, " ")
}
