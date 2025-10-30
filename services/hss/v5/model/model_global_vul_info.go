package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlobalVulInfo 漏洞对应的记录
type GlobalVulInfo struct {

	// 漏洞名称
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释** 漏洞ID **取值范围** 字符长度0-65535位
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 修复紧急度 **取值范围**: - immediate_repair：需尽快修复。 - delay_repair：可延后修复。 - not_needed_repair：暂可不修复。
	RepairNecessity *string `json:"repair_necessity,omitempty"`

	// **参数解释** 漏洞描述 **取值范围** 字符长度0-65535位
	Decription *string `json:"decription,omitempty"`

	// **参数解释** 解决方案 **取值范围** 字符长度0-65535位
	Solution *string `json:"solution,omitempty"`

	// **参数解释** URL链接 **取值范围** 字符长度0-65535位
	Url *string `json:"url,omitempty"`

	// **参数解释** 历史受影响镜像的个数 **取值范围** 取值0-65535
	HistoryNumber *int32 `json:"history_number,omitempty"`

	// **参数解释** 未处理镜像的格式 **取值范围** 取值0-65535
	UndealNumber *int32 `json:"undeal_number,omitempty"`

	// cve列表
	DataList *[]CveInfo `json:"data_list,omitempty"`
}

func (o GlobalVulInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalVulInfo struct{}"
	}

	return strings.Join([]string{"GlobalVulInfo", string(data)}, " ")
}
