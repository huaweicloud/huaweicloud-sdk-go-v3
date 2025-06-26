package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceInternalEndpointRequest Request Object
type ShowInstanceInternalEndpointRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 内网访问ID
	InternalEndpointsId string `json:"internal_endpoints_id"`
}

func (o ShowInstanceInternalEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceInternalEndpointRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceInternalEndpointRequest", string(data)}, " ")
}
