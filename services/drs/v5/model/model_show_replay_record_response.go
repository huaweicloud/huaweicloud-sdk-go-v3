package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReplayRecordResponse Response Object
type ShowReplayRecordResponse struct {

	// 列表中的项目总数，与分页无关。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 回放任务的SQL执行记录。
	SqlRecords     *[]ReplayExecuteRecords `json:"sql_records,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowReplayRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplayRecordResponse struct{}"
	}

	return strings.Join([]string{"ShowReplayRecordResponse", string(data)}, " ")
}
