package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronShowSubnetResponse Response Object
type NeutronShowSubnetResponse struct {
	Subnet         *NeutronSubnet `json:"subnet,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o NeutronShowSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowSubnetResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowSubnetResponse", string(data)}, " ")
}
