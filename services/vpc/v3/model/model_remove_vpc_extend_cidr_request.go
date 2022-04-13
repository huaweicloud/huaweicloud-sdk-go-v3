package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RemoveVpcExtendCidrRequest struct {
	// VPC资源ID

	VpcId string `json:"vpc_id"`

	Body *RemoveVpcExtendCidrRequestBody `json:"body,omitempty"`
}

func (o RemoveVpcExtendCidrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveVpcExtendCidrRequest struct{}"
	}

	return strings.Join([]string{"RemoveVpcExtendCidrRequest", string(data)}, " ")
}
