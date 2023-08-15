package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDeploymentRequest Request Object
type CreateDeploymentRequest struct {

	// 平台提供者，分别为hilens及ief。当为hilens时，请求部署在hilens平台的相关数据
	Provider *string `json:"provider,omitempty"`

	// 离线场景超期时间，单位分钟，范围在1-86400
	XExpiredTime *int32 `json:"X-Expired-Time,omitempty"`

	Body *DeploymentCreateRequest `json:"body,omitempty"`
}

func (o CreateDeploymentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentRequest struct{}"
	}

	return strings.Join([]string{"CreateDeploymentRequest", string(data)}, " ")
}
