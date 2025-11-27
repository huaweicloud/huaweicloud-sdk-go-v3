package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransitSubnetRequestBody
type CreateTransitSubnetRequestBody struct {
	TransitSubnet *CreateTransitSubnetOption `json:"transit_subnet"`
}

func (o CreateTransitSubnetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitSubnetRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTransitSubnetRequestBody", string(data)}, " ")
}
