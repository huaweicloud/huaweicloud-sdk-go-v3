package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulReportDataResponseInfoSumary 统计信息
type ShowVulReportDataResponseInfoSumary struct {

	// **参数解释**: 紧急修复漏洞数量 **取值范围**: 最小值0，最大值1000000
	VulNumRepairNecessity *int32 `json:"vul_num_repair_necessity,omitempty"`

	// **参数解释**: 未完成修复的漏洞数量 **取值范围**: 最小值0，最大值1000000
	VulNumUnfixed *int32 `json:"vul_num_unfixed,omitempty"`

	// **参数解释**: linxu漏洞数量 **取值范围**: 最小值0，最大值1000000
	VulNumLinux *int32 `json:"vul_num_linux,omitempty"`

	// **参数解释**: windows的漏洞数量 **取值范围**: 最小值0，最大值1000000
	VulNumWindows *int32 `json:"vul_num_windows,omitempty"`

	// **参数解释**: web-cms的漏洞数量 **取值范围**: 最小值0，最大值1000000
	VulNumWebCms *int32 `json:"vul_num_web_cms,omitempty"`

	// **参数解释**: 应用漏洞 **取值范围**: 最小值0，最大值1000000
	VulNumApp *int32 `json:"vul_num_app,omitempty"`

	// **参数解释**: 有风险的主机数量 **取值范围**: 最小值0，最大值1000000
	HostNumRisk *int32 `json:"host_num_risk,omitempty"`

	// **参数解释**: 有高危风险的主机数量 **取值范围**: 最小值0，最大值1000000
	HostNumHighRisk *int32 `json:"host_num_high_risk,omitempty"`

	// **参数解释**: 有中危风险的主机数量 **取值范围**: 最小值0，最大值1000000
	HostNumMediumRisk *int32 `json:"host_num_medium_risk,omitempty"`

	// **参数解释**: 有低危风险的主机数量 **取值范围**: 最小值0，最大值1000000
	HostNumLowRisk *int32 `json:"host_num_low_risk,omitempty"`

	// **参数解释**: 受影响的重要资产数量 **取值范围**: 最小值0，最大值1000000
	AffectAssetNumImportant *int32 `json:"affect_asset_num_important,omitempty"`

	// **参数解释**: 受影响的一般资产数量 **取值范围**: 最小值0，最大值1000000
	AffectAssetNumCommon *int32 `json:"affect_asset_num_common,omitempty"`

	// **参数解释**: 受影响的测试资产数量 **取值范围**: 最小值0，最大值1000000
	AffectAssetNumTest *int32 `json:"affect_asset_num_test,omitempty"`
}

func (o ShowVulReportDataResponseInfoSumary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulReportDataResponseInfoSumary struct{}"
	}

	return strings.Join([]string{"ShowVulReportDataResponseInfoSumary", string(data)}, " ")
}
