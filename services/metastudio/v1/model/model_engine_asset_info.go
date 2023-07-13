package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EngineAssetInfo struct {

	// 引擎资产的相对路径信息
	RelativePath *string `json:"relative_path,omitempty"`

	// 引擎资产的类型
	AssetType *string `json:"asset_type,omitempty"`
}

func (o EngineAssetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineAssetInfo struct{}"
	}

	return strings.Join([]string{"EngineAssetInfo", string(data)}, " ")
}
