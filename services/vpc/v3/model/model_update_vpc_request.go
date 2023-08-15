package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpcRequest Request Object
type UpdateVpcRequest struct {

	// VPC资源ID
	VpcId string `json:"vpc_id"`

	Body *UpdateVpcRequestBody `json:"body,omitempty"`
}

func (o UpdateVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcRequest struct{}"
	}

	return strings.Join([]string{"UpdateVpcRequest", string(data)}, " ")
}
