package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulScanPolicyResponse Response Object
type ShowVulScanPolicyResponse struct {

	// **参数解释**: 扫描周期 **取值范围**: - one_day：每天 - three_day：每三天 - one_week：每周 - one_month：每月
	ScanPeriod *string `json:"scan_period,omitempty"`

	// **参数解释**: \"扫描的漏洞类型列表\" **取值范围**: 最小值0，最大值5
	ScanVulTypes *[]string `json:"scan_vul_types,omitempty"`

	// **参数解释**: 扫描主机的范围 **取值范围**: - all_host：扫描全部主机 - specific_host：扫描指定主机
	ScanRangeType *string `json:"scan_range_type,omitempty"`

	// **参数解释**: 主机ID列表；当scan_range_type的值为specific_host时表示扫描的主机列表 **取值范围**: 最小值0，最大值20000
	HostIds *[]string `json:"host_ids,omitempty"`

	// **参数解释**: 可进行漏洞扫描的主机总数 **取值范围**: 最小值0，最大值20000
	TotalHostNum *int64 `json:"total_host_num,omitempty"`

	// **参数解释**: 扫描策略状态 **取值范围**: - open : 开启 - close : 关闭
	Status *string `json:"status,omitempty"`

	Time           *VulScanPolicyResponseInfoTime `json:"time,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowVulScanPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulScanPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowVulScanPolicyResponse", string(data)}, " ")
}
