package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulReportDataResponseInfoSumary 统计信息（该字段已废弃）
type ShowVulReportDataResponseInfoSumary struct {

	// **参数解释**: 紧急修复漏洞数量 **取值范围**: 最小值0，最大值1000000
	VulNumRepairNecessity *int32 `json:"vul_num_repair_necessity,omitempty"`
}

func (o ShowVulReportDataResponseInfoSumary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulReportDataResponseInfoSumary struct{}"
	}

	return strings.Join([]string{"ShowVulReportDataResponseInfoSumary", string(data)}, " ")
}
