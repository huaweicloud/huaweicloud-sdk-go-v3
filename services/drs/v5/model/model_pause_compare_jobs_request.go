package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PauseCompareJobsRequest Request Object
type PauseCompareJobsRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *OperateDataCompareJobReq `json:"body,omitempty"`
}

func (o PauseCompareJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseCompareJobsRequest struct{}"
	}

	return strings.Join([]string{"PauseCompareJobsRequest", string(data)}, " ")
}
