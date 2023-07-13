package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateSubnetRequestBody
type NeutronCreateSubnetRequestBody struct {
	Subnet *NeutronCreateSubnetOption `json:"subnet"`
}

func (o NeutronCreateSubnetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSubnetRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateSubnetRequestBody", string(data)}, " ")
}
