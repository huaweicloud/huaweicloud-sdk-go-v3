package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDeploymentNodesRequest Request Object
type AddDeploymentNodesRequest struct {

	// 部署ID
	DeploymentId string `json:"deployment_id"`

	// 平台提供者，分别为hilens及ief。当为hilens时，请求部署在hilens平台的相关数据。
	Provider *string `json:"provider,omitempty"`

	Body *DeploymentAddNodesRequest `json:"body,omitempty"`
}

func (o AddDeploymentNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDeploymentNodesRequest struct{}"
	}

	return strings.Join([]string{"AddDeploymentNodesRequest", string(data)}, " ")
}
