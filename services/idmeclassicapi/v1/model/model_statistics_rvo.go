package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticsRvo struct {

	// **参数解释：**  新增实例统计记录数，表示在指定时间区间内创建的实例数量。  **取值范围：**  不涉及。
	CreateCount *int32 `json:"createCount,omitempty"`

	// **参数解释：**  更新实例统计记录数，表示在指定时间区间内更新的实例数量。  **取值范围：**  不涉及。
	UpdateCount *int32 `json:"updateCount,omitempty"`

	// **参数解释：**  删除实例统计记录数，表示在指定时间区间内物理删除的实例数量。  **取值范围：**  不涉及。
	DeleteCount *int32 `json:"deleteCount,omitempty"`

	// **参数解释：**  软删除实例统计记录数，表示在指定时间区间内软删除的实例数量。  **取值范围：**  不涉及。
	LogicalDeleteCount *int32 `json:"logicalDeleteCount,omitempty"`
}

func (o StatisticsRvo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsRvo struct{}"
	}

	return strings.Join([]string{"StatisticsRvo", string(data)}, " ")
}
