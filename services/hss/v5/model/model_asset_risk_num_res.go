package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssetRiskNumRes **参数解释**: 资产风险数量 **取值范围**: 最小值0，最大值2147483647
type AssetRiskNumRes struct {
}

func (o AssetRiskNumRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetRiskNumRes struct{}"
	}

	return strings.Join([]string{"AssetRiskNumRes", string(data)}, " ")
}
