package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpcepServiceInfo struct {

	// 子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	Type *VpcServiceType `json:"type,omitempty"`

	// Vpc 端点地址
	VpcEndpointAddress *string `json:"vpc_endpoint_address,omitempty"`

	// Vpc 端点ID
	VpcEndpointId *string `json:"vpc_endpoint_id,omitempty"`

	// Vpc端点服务ID
	VpcEndpointServiceId *string `json:"vpc_endpoint_service_id,omitempty"`

	// UUID
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o UpdateVpcepServiceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcepServiceInfo struct{}"
	}

	return strings.Join([]string{"UpdateVpcepServiceInfo", string(data)}, " ")
}
