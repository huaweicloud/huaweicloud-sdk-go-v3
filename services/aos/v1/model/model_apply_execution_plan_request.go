package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ApplyExecutionPlanRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 用户希望操作的资源栈名
	StackName string `json:"stack_name"`

	// 执行计划的名字。如果未指定，则使用execution_plan_id作为execution_plan_name。
	ExecutionPlanName string `json:"execution_plan_name"`

	Body *ApplyExecutionPlanRequestBody `json:"body,omitempty"`
}

func (o ApplyExecutionPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyExecutionPlanRequest struct{}"
	}

	return strings.Join([]string{"ApplyExecutionPlanRequest", string(data)}, " ")
}
