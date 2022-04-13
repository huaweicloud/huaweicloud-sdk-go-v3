package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDeploymentRequest struct {
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`
	// 应用部署ID

	DeploymentId string `json:"deployment_id"`
	// 是否强制删除；默认为false。

	ForceDelete *string `json:"force_delete,omitempty"`
}

func (o DeleteDeploymentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentRequest", string(data)}, " ")
}
