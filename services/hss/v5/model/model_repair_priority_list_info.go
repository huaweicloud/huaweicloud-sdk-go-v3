package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepairPriorityListInfo struct {

	// **参数解释**: 修复优先级 **取值范围**: - Critical：紧急 - High：高 - Medium：中 - Low：低
	RepairPriority *string `json:"repair_priority,omitempty"`

	// **参数解释**: 当前修复优先级对应的主机数量 **取值范围**: 取值0-2147483647
	HostNum *int32 `json:"host_num,omitempty"`
}

func (o RepairPriorityListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepairPriorityListInfo struct{}"
	}

	return strings.Join([]string{"RepairPriorityListInfo", string(data)}, " ")
}
