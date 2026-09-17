package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableVacuumNumInfo struct {

	// **参数解释**： 等待的表数量。 **默认取值**： 不涉及。
	WaitingNum *int32 `json:"waiting_num,omitempty"`

	// **参数解释**： 运行中的表数量。 **默认取值**： 不涉及。
	RunningNum *int32 `json:"running_num,omitempty"`

	// **参数解释**： 已完成的表数量。 **默认取值**： 不涉及。
	FinishedNum *int32 `json:"finished_num,omitempty"`

	// **参数解释**： 取消的表数量。 **默认取值**： 不涉及。
	CanceledNum *int32 `json:"canceled_num,omitempty"`

	// **参数解释**： 总数。 **默认取值**： 不涉及。
	TotalNum *int32 `json:"total_num,omitempty"`
}

func (o TableVacuumNumInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableVacuumNumInfo struct{}"
	}

	return strings.Join([]string{"TableVacuumNumInfo", string(data)}, " ")
}
