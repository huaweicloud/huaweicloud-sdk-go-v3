package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerIpsPageInfo struct {

	// **参数解释**： 查询返回记录的数量限制 **约束限制**：   不涉及 **取值范围**： 1-1024 **默认取值**：   不涉及
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量，表示查询该偏移量后面的记录 **约束限制**：   不涉及 **取值范围**： 0 - 1024 **默认取值**：   不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 自定义IPS规则列表 **约束限制**：   不涉及 **取值范围**： 不涉及 **默认取值**：   不涉及
	Records *[]CustomerIpsListVo `json:"records,omitempty"`

	// **参数解释**： 自定义IPS规则数量 **约束限制**：   不涉及 **取值范围**： 不涉及 **默认取值**：   不涉及
	Total *int32 `json:"total,omitempty"`
}

func (o CustomerIpsPageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerIpsPageInfo struct{}"
	}

	return strings.Join([]string{"CustomerIpsPageInfo", string(data)}, " ")
}
