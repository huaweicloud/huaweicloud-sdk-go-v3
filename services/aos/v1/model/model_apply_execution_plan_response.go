package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ApplyExecutionPlanResponse struct {

	// 部署生成的唯一ID，由资源编排服务生成
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
