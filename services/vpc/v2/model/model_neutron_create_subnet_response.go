package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateSubnetResponse Response Object
type NeutronCreateSubnetResponse struct {
	Subnet         *NeutronSubnet `json:"subnet,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o NeutronCreateSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSubnetResponse struct{}"
	}

	return strings.Join([]string{"NeutronCreateSubnetResponse", string(data)}, " ")
}
