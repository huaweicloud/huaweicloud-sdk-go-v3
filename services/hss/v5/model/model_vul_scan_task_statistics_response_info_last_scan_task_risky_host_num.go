package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulScanTaskStatisticsResponseInfoLastScanTaskRiskyHostNum **参数解释**: 影响的主机数 **取值范围**: 不涉及
type VulScanTaskStatisticsResponseInfoLastScanTaskRiskyHostNum struct {

	// **参数解释**: 应急漏洞影响的主机数 **取值范围**: 最小值0，最大值2147483647
	UrgentVul *int32 `json:"urgent_vul,omitempty"`
}

func (o VulScanTaskStatisticsResponseInfoLastScanTaskRiskyHostNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulScanTaskStatisticsResponseInfoLastScanTaskRiskyHostNum struct{}"
	}

	return strings.Join([]string{"VulScanTaskStatisticsResponseInfoLastScanTaskRiskyHostNum", string(data)}, " ")
}
