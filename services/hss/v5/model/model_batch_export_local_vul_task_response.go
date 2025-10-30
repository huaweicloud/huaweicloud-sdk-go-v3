package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchExportLocalVulTaskResponse Response Object
type BatchExportLocalVulTaskResponse struct {

	// **参数解释**: 导出记录总数 **取值范围**: 最小值0，最大值2147483647
	RecordTotalNum *int64 `json:"record_total_num,omitempty"`

	// **参数解释**: 导出任务ID **取值范围**: 字符长度1-256位
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchExportLocalVulTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExportLocalVulTaskResponse struct{}"
	}

	return strings.Join([]string{"BatchExportLocalVulTaskResponse", string(data)}, " ")
}
