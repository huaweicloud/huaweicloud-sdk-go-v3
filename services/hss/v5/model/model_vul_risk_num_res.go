package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulRiskNumRes **参数解释**: 漏洞数 **取值范围**: 最小值0，最大值2147483647
type VulRiskNumRes struct {
}

func (o VulRiskNumRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulRiskNumRes struct{}"
	}

	return strings.Join([]string{"VulRiskNumRes", string(data)}, " ")
}
