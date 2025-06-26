package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceInternalEndpointRequest Request Object
type DeleteInstanceInternalEndpointRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 内网访问ID
	InternalEndpointsId string `json:"internal_endpoints_id"`
}

func (o DeleteInstanceInternalEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceInternalEndpointRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceInternalEndpointRequest", string(data)}, " ")
}
