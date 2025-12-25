package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpcEndpointServiceRequestBody struct {

	// 更新VpcepServiceInfo
	List *[]UpdateVpcepServiceInfo `json:"list,omitempty"`

	// 子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// UUID
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o UpdateVpcEndpointServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcEndpointServiceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVpcEndpointServiceRequestBody", string(data)}, " ")
}
