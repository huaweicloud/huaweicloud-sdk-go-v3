package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CriticalVulNumRes **参数解释**: 紧急漏洞数量 **取值范围**: 最小值0，最大值2147483647
type CriticalVulNumRes struct {
}

func (o CriticalVulNumRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CriticalVulNumRes struct{}"
	}

	return strings.Join([]string{"CriticalVulNumRes", string(data)}, " ")
}
