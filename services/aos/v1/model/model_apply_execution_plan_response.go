package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ApplyExecutionPlanResponse struct {

	// 标识部署的唯一Id，此Id由资源编排服务在触发部署、回滚等操作时生成，为UUID。
	DeploymentId   *string `json:"deployment_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ApplyExecutionPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyExecutionPlanResponse struct{}"
	}

	return strings.Join([]string{"ApplyExecutionPlanResponse", string(data)}, " ")
}
