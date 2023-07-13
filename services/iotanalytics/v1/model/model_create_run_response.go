package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRunResponse Response Object
type CreateRunResponse struct {

	// 作业运行ID。
	RunId *string `json:"run_id,omitempty"`

	// 作业ID。
	JobId *string `json:"job_id,omitempty"`

	// 作业类型。
	JobType *string `json:"job_type,omitempty"`

	// 创建运行时间。
	CreatedTime *string `json:"created_time,omitempty"`

	SqlJob         *SqlJobRunResponseBody `json:"sql_job,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CreateRunResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRunResponse struct{}"
	}

	return strings.Join([]string{"CreateRunResponse", string(data)}, " ")
}
