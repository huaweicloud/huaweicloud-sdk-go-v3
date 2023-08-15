package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpcPeeringResponse Response Object
type DeleteVpcPeeringResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcPeeringResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcPeeringResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpcPeeringResponse", string(data)}, " ")
}
