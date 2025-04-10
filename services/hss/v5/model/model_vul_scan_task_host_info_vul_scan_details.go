package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulScanTaskHostInfoVulScanDetails struct {

	// 扫描的漏洞类型，包含如下：   - linux_vul : linux漏洞   - windows_vul : windows漏洞   - web_cms : Web-CMS漏洞   - app_vul : 应用漏洞   - urgent_vul : 应急漏洞
	VulType *string `json:"vul_type,omitempty"`

	// 该类型漏洞的扫描状态，包含如下：   - scanning : 扫描中   - success : 扫描成功   - failed : 扫描失败
	Status *string `json:"status,omitempty"`

	// 扫描失败的原因，只有扫描失败的漏洞类型有该字段
	FailedReason *string `json:"failed_reason,omitempty"`

	// 扫描的漏洞列表，只有漏洞类型为应急漏洞有该字段
	ScanVulList *[]VulScanTaskHostInfoScanVulList `json:"scan_vul_list,omitempty"`
}

func (o VulScanTaskHostInfoVulScanDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulScanTaskHostInfoVulScanDetails struct{}"
	}

	return strings.Join([]string{"VulScanTaskHostInfoVulScanDetails", string(data)}, " ")
}
