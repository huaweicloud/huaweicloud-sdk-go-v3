package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShortJobResponse Response Object
type ShowShortJobResponse struct {
	JobType *ShortJobType `json:"job_type,omitempty"`

	// 任务id。
	JobId *string `json:"job_id,omitempty"`

	State *JobState `json:"state,omitempty"`

	// 当任务失败时呈现,失败错误码。
	JobFailedCode *string `json:"job_failed_code,omitempty"`

	// 当任务失败时呈现,失败原因。
	JobFailedReason *string `json:"job_failed_reason,omitempty"`

	// 任务创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 任务状态更新时间。
	LastupdateTime *int64 `json:"lastupdate_time,omitempty"`

	AssessResult   *AssessResult `json:"assess_result,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowShortJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShortJobResponse struct{}"
	}

	return strings.Join([]string{"ShowShortJobResponse", string(data)}, " ")
}
