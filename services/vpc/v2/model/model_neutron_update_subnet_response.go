package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateSubnetResponse Response Object
type NeutronUpdateSubnetResponse struct {
	Subnet         *NeutronSubnet `json:"subnet,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o NeutronUpdateSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateSubnetResponse struct{}"
	}

	return strings.Join([]string{"NeutronUpdateSubnetResponse", string(data)}, " ")
}
