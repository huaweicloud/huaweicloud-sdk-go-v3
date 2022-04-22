package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HistoryRecord struct {

	// 构建记录id--唯一key
	RecordId *string `json:"record_id,omitempty"`

	// 任务id
	JobId *string `json:"job_id,omitempty"`

	// 构建编号
	BuildNumber *int32 `json:"build_number,omitempty"`

	// 构建开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 构建结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 构建结果
	Result *string `json:"result,omitempty"`
}

func (o HistoryRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryRecord struct{}"
	}

	return strings.Join([]string{"HistoryRecord", string(data)}, " ")
}
