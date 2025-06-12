package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobUpdateRecordListVoResult 结果
type JobUpdateRecordListVoResult struct {

	// job_update_records
	JobUpdateRecords *[]JobUpdateRecordListVoResultJobUpdateRecords `json:"job_update_records,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`
}

func (o JobUpdateRecordListVoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobUpdateRecordListVoResult struct{}"
	}

	return strings.Join([]string{"JobUpdateRecordListVoResult", string(data)}, " ")
}
