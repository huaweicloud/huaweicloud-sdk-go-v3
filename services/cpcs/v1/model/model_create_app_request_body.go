package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAppRequestBody struct {

	// 应用名称
	AppName string `json:"app_name"`

	// 应用所属的VPC的ID
	VpcId string `json:"vpc_id"`

	// 应用所属的VPC的名称
	VpcName string `json:"vpc_name"`

	// 应用所属的VPC的子网ID
	SubnetId string `json:"subnet_id"`

	// 应用所属的VPC的子网名称
	SubnetName string `json:"subnet_name"`

	// 应用描述
	Description *string `json:"description,omitempty"`
}

func (o CreateAppRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAppRequestBody", string(data)}, " ")
}
