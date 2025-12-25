package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpcEndpointServiceRequest Request Object
type UpdateVpcEndpointServiceRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 虚拟私有云ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	Body *UpdateVpcEndpointServiceRequestBody `json:"body,omitempty"`
}

func (o UpdateVpcEndpointServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcEndpointServiceRequest struct{}"
	}

	return strings.Join([]string{"UpdateVpcEndpointServiceRequest", string(data)}, " ")
}
