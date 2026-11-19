package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulTotalNumRes **参数解释**: 漏洞总数 **取值范围**: 最小值0，最大值2147483647
type VulTotalNumRes struct {
}

func (o VulTotalNumRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulTotalNumRes struct{}"
	}

	return strings.Join([]string{"VulTotalNumRes", string(data)}, " ")
}
