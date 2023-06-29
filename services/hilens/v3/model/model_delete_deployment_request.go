package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeploymentRequest Request Object
type DeleteDeploymentRequest struct {

	// 部署ID
	DeploymentId string `json:"deployment_id"`

	// 强制删除，为true时表示前置删除
	ForceDelete *bool `json:"force_delete,omitempty"`

	// 平台提供者，分别为hilens及ief。当为hilens时，请求部署在hilens平台的相关数据
	Provider *string `json:"provider,omitempty"`

	// 离线场景超期时间，单位分钟，范围在1-86400
	XExpiredTime *int32 `json:"X-Expired-Time,omitempty"`
}

func (o DeleteDeploymentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentRequest", string(data)}, " ")
}
