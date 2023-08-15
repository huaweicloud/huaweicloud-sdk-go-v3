package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateSubnetRequestBody
type NeutronUpdateSubnetRequestBody struct {
	Subnet *NeutronUpdateSubnetOption `json:"subnet"`
}

func (o NeutronUpdateSubnetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateSubnetRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronUpdateSubnetRequestBody", string(data)}, " ")
}
