package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StageRecord 作业流程阶段记录。
type StageRecord struct {

	// **参数解释**：阶段记录序号，顺序递增，最大序号记录为当前最新记录。  **取值范围**：不涉及。
	RecordOrder *int32 `json:"record_order,omitempty"`

	// **参数解释**：主阶段信息列表。
	Stages *[]StageInfoWithSub `json:"stages,omitempty"`
}

func (o StageRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StageRecord struct{}"
	}

	return strings.Join([]string{"StageRecord", string(data)}, " ")
}
