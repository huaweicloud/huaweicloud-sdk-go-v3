package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchStartOrStopTasksResponse struct {

	// 成功的个数
	SuccessCount *int32 `json:"success_count,omitempty" xml:"success_count"`

	// 失败的个数
	FailureCount *int32 `json:"failure_count,omitempty" xml:"failure_count"`

	// 失败的详情
	Failure *[]TaskBeanFacade `json:"failure,omitempty" xml:"failure"`

	// 成功的任务信息
	Success        *[]TaskBeanFacade `json:"success,omitempty" xml:"success"`
	HttpStatusCode int               `json:"-"`
}

func (o BatchStartOrStopTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartOrStopTasksResponse struct{}"
	}

	return strings.Join([]string{"BatchStartOrStopTasksResponse", string(data)}, " ")
}
