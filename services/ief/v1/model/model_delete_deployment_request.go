package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeploymentRequest Request Object
type DeleteDeploymentRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 应用部署ID
	DeploymentId string `json:"deployment_id"`

	// 是否强制删除。默认为false。 如果强制删除，则会忽略边缘节点是否有残留应用，直接删除云端应用。 如果不强制删除，则会等待边缘节点的应用删除成功后，再删除云端应用。
	ForceDelete *string `json:"force_delete,omitempty"`
}

func (o DeleteDeploymentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentRequest", string(data)}, " ")
}
