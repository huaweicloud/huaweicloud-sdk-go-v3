package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExecutionPlanResponse Response Object
type CreateExecutionPlanResponse struct {

	// 执行计划（execution_plan）的唯一Id。  此Id由资源编排服务在生成执行计划的时候生成，为UUID。  由于执行计划名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的执行计划，删除，再重新创建一个同名执行计划。  对于团队并行开发，用户可能希望确保，当前我操作的执行计划就是我认为的那个，而不是其他队友删除后创建的同名执行计划。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的执行计划所对应的ID都不相同，更新不会影响ID。如果给与的execution_plan_id和当前执行计划的ID不一致，则返回400
	ExecutionPlanId *string `json:"execution_plan_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreateExecutionPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExecutionPlanResponse struct{}"
	}

	return strings.Join([]string{"CreateExecutionPlanResponse", string(data)}, " ")
}
