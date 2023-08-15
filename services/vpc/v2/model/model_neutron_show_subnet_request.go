package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronShowSubnetRequest Request Object
type NeutronShowSubnetRequest struct {

	// 子网ID
	SubnetId string `json:"subnet_id"`
}

func (o NeutronShowSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowSubnetRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowSubnetRequest", string(data)}, " ")
}
