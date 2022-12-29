package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDeploymentRequest struct {

	// 应用部署ID
	DeploymentId string `json:"deployment_id"`

	// 平台提供者，分别为hilens及ief。当为hilens时，请求部署在hilens平台的相关数据。
	Provider *string `json:"provider,omitempty"`
}

func (o ShowDeploymentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentRequest struct{}"
	}

	return strings.Join([]string{"ShowDeploymentRequest", string(data)}, " ")
}
