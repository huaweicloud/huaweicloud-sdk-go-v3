package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartCompareJobsRequest Request Object
type RestartCompareJobsRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *OperateDataCompareJobReq `json:"body,omitempty"`
}

func (o RestartCompareJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartCompareJobsRequest struct{}"
	}

	return strings.Join([]string{"RestartCompareJobsRequest", string(data)}, " ")
}
