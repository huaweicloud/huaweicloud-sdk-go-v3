package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulScanTaskHostInfoScanVulList struct {

	// 漏洞id
	VulId *string `json:"vul_id,omitempty"`

	// 漏洞名称
	VulName *string `json:"vul_name,omitempty"`

	// 该漏洞的扫描状态，包含如下：   - scanning : 扫描中   - success : 扫描成功   - failed : 扫描失败
	Status *string `json:"status,omitempty"`

	// 该漏洞扫描失败的原因，只有扫描失败的漏洞有该字段
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o VulScanTaskHostInfoScanVulList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulScanTaskHostInfoScanVulList struct{}"
	}

	return strings.Join([]string{"VulScanTaskHostInfoScanVulList", string(data)}, " ")
}
