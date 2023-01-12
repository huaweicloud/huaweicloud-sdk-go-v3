package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJob2Response struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 任务状态。  - pending:等待中 - running:运行中 - success:成功 - failed:失败
	Status *string `json:"status,omitempty"`

	// 任务类型。
	JobType *string `json:"job_type,omitempty"`

	// 任务名称。
	JobName *string `json:"job_name,omitempty"`

	// 关联图名称。
	RelatedGraph *string `json:"related_graph,omitempty"`

	// 任务开始时间，格式为UTC,\"yyyy-MM-dd'T'HH:mm:ss\"
	BeginTime *string `json:"begin_time,omitempty"`

	// 任务结束时间，格式为UTC,\"yyyy-MM-dd'T'HH:mm:ss\"
	EndTime *string `json:"end_time,omitempty"`

	JobDetail *ShowJobRespJobDetail `json:"job_detail,omitempty"`

	// 任务失败原因
	FailReason *string `json:"fail_reason,omitempty"`

	// 任务执行进度，预留字段，暂未使用。
	JobProgress    *float64 `json:"job_progress,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowJob2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJob2Response struct{}"
	}

	return strings.Join([]string{"ShowJob2Response", string(data)}, " ")
}
