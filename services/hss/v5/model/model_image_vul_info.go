package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageVulInfo struct {

	// **参数解释**: 漏洞ID **取值范围**: 字符长度1-64位
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 漏洞名称 **取值范围**: 字符长度0-256位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**: 修复紧急度 **取值范围**: - immediate_repair：紧急修复 - delay_repair：延迟修复 - not_needed_repair：无需修复
	RepairNecessity *string `json:"repair_necessity,omitempty"`

	// **参数解释**: 漏洞描述 **取值范围**: 字符长度0-4096位
	Description *string `json:"description,omitempty"`

	// **参数解释**: 漏洞位置 **取值范围**: 字符长度0-512位
	Position *string `json:"position,omitempty"`

	// **参数解释**: 软件名称 **取值范围**: 字符长度0-128位
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 软件路径 **取值范围**: 字符长度0-512位
	AppPath *string `json:"app_path,omitempty"`

	// **参数解释**: 软件版本 **取值范围**: 字符长度0-128位
	Version *string `json:"version,omitempty"`

	// **参数解释**: 修复建议 **取值范围**: 字符长度0-4096位
	Solution *string `json:"solution,omitempty"`

	// **参数解释**: URL **取值范围**: 字符长度0-2048位
	Url *string `json:"url,omitempty"`

	// **参数解释**: CVE列表 **取值范围**: 取值0-100个CveInfo对象
	DataList *[]CveInfo `json:"data_list,omitempty"`

	// **参数解释**: 漏洞状态 **取值范围**: 字符长度0-64位
	Status *string `json:"status,omitempty"`

	// **参数解释**: 危险程度 **取值范围**: - Low：低危 - Medium：中危 - High：高危 - Critical：危急
	SeverityLevel *string `json:"severity_level,omitempty"`
}

func (o ImageVulInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageVulInfo struct{}"
	}

	return strings.Join([]string{"ImageVulInfo", string(data)}, " ")
}
