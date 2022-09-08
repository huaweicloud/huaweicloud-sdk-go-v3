package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchStartOrStopTasksResponse struct {

	// 成功的个数
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 失败的个数
	FailureCount *int32 `json:"failure_count,omitempty"`

	// 失败的详情
	Failure *[]TaskBeanFacade `json:"failure,omitempty"`

	// 成功的任务信息
	Success        *[]TaskBeanFacade `json:"success,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o BatchStartOrStopTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartOrStopTasksResponse struct{}"
	}

	return strings.Join([]string{"BatchStartOrStopTasksResponse", string(data)}, " ")
}
