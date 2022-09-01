package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDatasetResponse struct {

	// 作业结果总个数。
	Count *int64 `json:"count,omitempty" xml:"count"`

	// 作业运行ID。
	RunId *string `json:"run_id,omitempty" xml:"run_id"`

	// 作业类型。
	JobType *string `json:"job_type,omitempty" xml:"job_type"`

	SqlJob         *SqlJobQueryDataset `json:"sql_job,omitempty" xml:"sql_job"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowDatasetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatasetResponse struct{}"
	}

	return strings.Join([]string{"ShowDatasetResponse", string(data)}, " ")
}
