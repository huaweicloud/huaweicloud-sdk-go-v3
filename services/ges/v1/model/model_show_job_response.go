package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobResponse struct {

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage"`

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode"`

	// 任务ID。
	JobId *string `json:"jobId,omitempty" xml:"jobId"`

	// 任务状态。  - pending:等待中 - running:运行中 - success:成功 - failed:失败
	Status *string `json:"status,omitempty" xml:"status"`

	// 任务类型。
	JobType *string `json:"jobType,omitempty" xml:"jobType"`

	// 任务名称。
	JobName *string `json:"jobName,omitempty" xml:"jobName"`

	// 关联图名称。
	RelatedGraph *string `json:"relatedGraph,omitempty" xml:"relatedGraph"`

	// 任务开始时间，格式为UTC,\"yyyy-MM-dd'T'HH:mm:ss\"
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime"`

	// 任务结束时间，格式为UTC,\"yyyy-MM-dd'T'HH:mm:ss\"
	EndTime *string `json:"endTime,omitempty" xml:"endTime"`

	JobDetail *JobDetail `json:"jobDetail,omitempty" xml:"jobDetail"`

	// 任务失败原因
	FailReason *string `json:"failReason,omitempty" xml:"failReason"`

	// 任务执行进度，预留字段，暂未使用。
	JobProgress    *float64 `json:"jobProgress,omitempty" xml:"jobProgress"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResponse struct{}"
	}

	return strings.Join([]string{"ShowJobResponse", string(data)}, " ")
}
