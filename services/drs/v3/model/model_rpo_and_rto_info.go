package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RpoAndRtoInfo RPO和RTO信息体
type RpoAndRtoInfo struct {

	// **参数解释**：  数据复制的检查点位置，标识已同步数据的Binlog位置。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	CheckPoint *string `json:"check_point,omitempty"`

	// **参数解释**：  数据同步延迟时间。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Delay *string `json:"delay,omitempty"`

	// **参数解释**：  全局事务ID集合（Global Transaction Identifier），用于标识已执行的事务。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	GtidSet *string `json:"gtid_set,omitempty"`

	// **参数解释**：  RPO或RTO的统计时间点，格式为\"yyyy-MM-dd HH:mm:ss\"。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Time *string `json:"time,omitempty"`
}

func (o RpoAndRtoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RpoAndRtoInfo struct{}"
	}

	return strings.Join([]string{"RpoAndRtoInfo", string(data)}, " ")
}
