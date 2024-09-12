package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OutputAssetConfig 输出视频资产信息配置。
type OutputAssetConfig struct {

	// **参数解释**： 输出视频资产名称。 **约束限制**： 不涉及。 **取值范围**： 字符长度0-256位。 **默认取值**： 不涉及。
	AssetName string `json:"asset_name"`
}

func (o OutputAssetConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputAssetConfig struct{}"
	}

	return strings.Join([]string{"OutputAssetConfig", string(data)}, " ")
}
