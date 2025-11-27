package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTransitSubnetRequestBody
type UpdateTransitSubnetRequestBody struct {
	TransitSubnet *UpdateTransitSubnetOption `json:"transit_subnet"`
}

func (o UpdateTransitSubnetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransitSubnetRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTransitSubnetRequestBody", string(data)}, " ")
}
