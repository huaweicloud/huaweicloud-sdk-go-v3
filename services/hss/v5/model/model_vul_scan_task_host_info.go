package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulScanTaskHostInfo struct {

	// **参数解释**: 主机ID **取值范围**: 字符长度1-128位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 主机名称 **取值范围**: 字符长度0-128位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 弹性公网IP地址 **取值范围**: 字符长度0-128位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 私有IP地址 **取值范围**: 字符长度0-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 主机的资产重要性 **取值范围**: - important：重要资产 - common：一般资产 - test：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 主机的扫描状态 **取值范围**: - scanning：扫描中 - success：扫描成功 - failed：扫描失败
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释**: 漏洞扫描失败的原因列表（即将废弃，建议使用“vul_scan_details”字段） **取值范围**: 最小值0，最大值2147483647
	FailedReasons *[]VulScanTaskHostInfoFailedReasons `json:"failed_reasons,omitempty"`

	// **参数解释**: 该主机的扫描详情信息列表 **取值范围**: 最小值0，最大值2147483647
	VulScanDetails *[]VulScanTaskHostInfoVulScanDetails `json:"vul_scan_details,omitempty"`
}

func (o VulScanTaskHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulScanTaskHostInfo struct{}"
	}

	return strings.Join([]string{"VulScanTaskHostInfo", string(data)}, " ")
}
