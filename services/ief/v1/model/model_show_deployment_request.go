package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDeploymentRequest struct {
	// 应用部署ID

	DeploymentId string `json:"deployment_id"`
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o ShowDeploymentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentRequest struct{}"
	}

	return strings.Join([]string{"ShowDeploymentRequest", string(data)}, " ")
}
