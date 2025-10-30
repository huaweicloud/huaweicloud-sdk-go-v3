package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulRepairFailedDetailInfo **参数解释**: 漏洞修复失败原因 **取值范围**: 不涉及
type VulRepairFailedDetailInfo struct {

	// **参数解释**: 软件名称 **取值范围**: 字符长度1-512位
	Software *string `json:"software,omitempty"`

	// **参数解释**: 漏洞修复失败原因详情 **取值范围**: 字符长度0-65535位
	Reason *string `json:"reason,omitempty"`

	// **参数解释**: 漏洞修复失败原因解释说明 **取值范围**: 字符长度0-256位
	ReasonDescribtion *string `json:"reason_describtion,omitempty"`

	// **参数解释**: 解决方式说明 **取值范围**: 字符长度0-65535位
	ReasonSolution *string `json:"reason_solution,omitempty"`
}

func (o VulRepairFailedDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulRepairFailedDetailInfo struct{}"
	}

	return strings.Join([]string{"VulRepairFailedDetailInfo", string(data)}, " ")
}
