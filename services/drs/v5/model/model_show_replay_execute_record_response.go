package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReplayExecuteRecordResponse Response Object
type ShowReplayExecuteRecordResponse struct {

	// 列表中的项目总数，与分页无关。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 回放任务的SQL执行记录。
	SqlRecords     *[]ReplayExecuteRecords `json:"sql_records,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowReplayExecuteRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplayExecuteRecordResponse struct{}"
	}

	return strings.Join([]string{"ShowReplayExecuteRecordResponse", string(data)}, " ")
}
