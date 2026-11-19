package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaselineRiskNumRes **参数解释**: 基线风险数量 **取值范围**: 最小值0，最大值2147483647
type BaselineRiskNumRes struct {
}

func (o BaselineRiskNumRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaselineRiskNumRes struct{}"
	}

	return strings.Join([]string{"BaselineRiskNumRes", string(data)}, " ")
}
