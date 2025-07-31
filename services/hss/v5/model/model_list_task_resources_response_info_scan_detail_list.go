package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskResourcesResponseInfoScanDetailList 扫描项信息
type ListTaskResourcesResponseInfoScanDetailList struct {

	// 扫描项类型，包含如下   - cluster_vul：集群漏洞   - risk_assessment：风险评估   - benchmark：安全合规
	ScanType *string `json:"scan_type,omitempty"`

	// 扫描状态，包含如下：   - scanning：扫描中   - success：扫描成功   - failed：扫描失败
	Status *string `json:"status,omitempty"`

	// 扫描失败的原因
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o ListTaskResourcesResponseInfoScanDetailList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskResourcesResponseInfoScanDetailList struct{}"
	}

	return strings.Join([]string{"ListTaskResourcesResponseInfoScanDetailList", string(data)}, " ")
}
