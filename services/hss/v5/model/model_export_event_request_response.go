package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportEventRequestResponse Response Object
type ExportEventRequestResponse struct {

	// **参数解释**： 导出记录总数 **取值范围**： 不涉及
	RecordTotalNum *int32 `json:"record_total_num,omitempty"`

	// **参数解释**： 导出任务ID **取值范围**： 不涉及
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportEventRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportEventRequestResponse struct{}"
	}

	return strings.Join([]string{"ExportEventRequestResponse", string(data)}, " ")
}
