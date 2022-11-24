package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobResponse struct {

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。
	ErrorCode *string `json:"errorCode,omitempty"`

	// 任务ID。
	JobId *string `json:"jobId,omitempty"`

	// 任务状态。  - pending:等待中 - running:运行中 - success:成功 - failed:失败
	Status *string `json:"status,omitempty"`

	// 任务类型。
	JobType *string `json:"jobType,omitempty"`

	// 任务名称。
	JobName *string `json:"jobName,omitempty"`

	// 关联图名称。
	RelatedGraph *string `json:"relatedGraph,omitempty"`

	// 任务开始时间，格式为UTC,\"yyyy-MM-dd'T'HH:mm:ss\"
	BeginTime *string `json:"beginTime,omitempty"`

	// 任务结束时间，格式为UTC,\"yyyy-MM-dd'T'HH:mm:ss\"
	EndTime *string `json:"endTime,omitempty"`

	JobDetail *JobDetail `json:"jobDetail,omitempty"`

	// 任务失败原因
	FailReason *string `json:"failReason,omitempty"`

	// 任务执行进度，预留字段，暂未使用。
	JobProgress    *float64 `json:"jobProgress,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResponse struct{}"
	}

	return strings.Join([]string{"ShowJobResponse", string(data)}, " ")
}
