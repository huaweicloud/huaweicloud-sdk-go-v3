package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronDeleteSubnetResponse Response Object
type NeutronDeleteSubnetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteSubnetResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteSubnetResponse", string(data)}, " ")
}
