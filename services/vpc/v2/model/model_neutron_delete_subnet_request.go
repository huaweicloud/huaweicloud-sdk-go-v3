package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronDeleteSubnetRequest Request Object
type NeutronDeleteSubnetRequest struct {

	// 子网ID
	SubnetId string `json:"subnet_id"`
}

func (o NeutronDeleteSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteSubnetRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteSubnetRequest", string(data)}, " ")
}
