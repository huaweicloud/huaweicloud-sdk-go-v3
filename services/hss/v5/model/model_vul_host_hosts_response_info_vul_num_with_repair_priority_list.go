package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulHostHostsResponseInfoVulNumWithRepairPriorityList struct {

	// **参数解释**: 漏洞修复优先级 **取值范围**: 字符长度0-64位
	RepairPriority *string `json:"repair_priority,omitempty"`

	// **参数解释**: 该优先级下的漏洞数量 **取值范围**: 最小值0，最大值2147483647
	VulNum *int32 `json:"vul_num,omitempty"`
}

func (o VulHostHostsResponseInfoVulNumWithRepairPriorityList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulHostHostsResponseInfoVulNumWithRepairPriorityList struct{}"
	}

	return strings.Join([]string{"VulHostHostsResponseInfoVulNumWithRepairPriorityList", string(data)}, " ")
}
