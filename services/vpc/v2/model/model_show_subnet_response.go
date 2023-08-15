package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubnetResponse Response Object
type ShowSubnetResponse struct {
	Subnet         *Subnet `json:"subnet,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubnetResponse struct{}"
	}

	return strings.Join([]string{"ShowSubnetResponse", string(data)}, " ")
}
