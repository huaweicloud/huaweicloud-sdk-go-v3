package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulScanTaskHostInfoScanVulList struct {

	// **参数解释**: 漏洞id **取值范围**: 字符长度0-256位
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 漏洞名称 **取值范围**: 字符长度0-1024位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**: 该漏洞的扫描状态 **取值范围**: - scanning：扫描中 - success：扫描成功 - failed：扫描失败
	Status *string `json:"status,omitempty"`

	// **参数解释**: 该漏洞扫描失败的原因，只有扫描失败的漏洞有该字段 **取值范围**: 字符长度0-1024位
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o VulScanTaskHostInfoScanVulList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulScanTaskHostInfoScanVulList struct{}"
	}

	return strings.Join([]string{"VulScanTaskHostInfoScanVulList", string(data)}, " ")
}
