package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpcPeeringRequestBody
type CreateVpcPeeringRequestBody struct {
	Peering *CreateVpcPeeringOption `json:"peering"`
}

func (o CreateVpcPeeringRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcPeeringRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVpcPeeringRequestBody", string(data)}, " ")
}
