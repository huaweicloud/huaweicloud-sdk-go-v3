package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticsRvo struct {

	// **参数解释：**  新增统计记录数。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateCount *int32 `json:"createCount,omitempty"`

	// **参数解释：**  删除统计记录数。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	DeleteCount *int32 `json:"deleteCount,omitempty"`

	// **参数解释：**  软删除统计记录数。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LogicalDeleteCount *int32 `json:"logicalDeleteCount,omitempty"`

	// **参数解释：**  更新统计记录数。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	UpdateCount *int32 `json:"updateCount,omitempty"`
}

func (o StatisticsRvo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsRvo struct{}"
	}

	return strings.Join([]string{"StatisticsRvo", string(data)}, " ")
}
