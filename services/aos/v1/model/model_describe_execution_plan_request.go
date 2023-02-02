package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DescribeExecutionPlanRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 用户希望操作的资源栈名称
	StackName string `json:"stack_name"`

	// 执行计划的名字。
	ExecutionPlanName string `json:"execution_plan_name"`

	// 用户希望描述的资源栈的Id。若stack_name和stack_id同时存在，则资源编排服务会检查是否两个匹配，否则返回400
	StackId *string `json:"stack_id,omitempty"`

	// 执行计划ID(uuid)
	ExecutionPlanId *string `json:"execution_plan_id,omitempty"`
}

func (o DescribeExecutionPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeExecutionPlanRequest struct{}"
	}

	return strings.Join([]string{"DescribeExecutionPlanRequest", string(data)}, " ")
}
