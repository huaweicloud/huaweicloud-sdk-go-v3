package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeVpcRequestBody This is a auto create Body Object
type ChangeVpcRequestBody struct {

	// 虚拟机私有云ID
	VpcId string `json:"vpc_id"`

	Nic *ChangeVpcNicBody `json:"nic"`
}

func (o ChangeVpcRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVpcRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeVpcRequestBody", string(data)}, " ")
}
