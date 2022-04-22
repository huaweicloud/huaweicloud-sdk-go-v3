package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListExecuteJobResponse struct {

	// 作业列表总数。
	TotalRecord *int32 `json:"totalRecord,omitempty"`

	// 作业列表。
	JobExecutions  *[]JobExeResult `json:"job_executions,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListExecuteJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExecuteJobResponse struct{}"
	}

	return strings.Join([]string{"ListExecuteJobResponse", string(data)}, " ")
}
