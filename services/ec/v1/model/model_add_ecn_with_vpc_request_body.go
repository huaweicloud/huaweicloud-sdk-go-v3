package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddEcnWithVpcRequestBody struct {

	// 虚拟私有云ID
	VpcId string `json:"vpc_id"`

	// 子网ID
	SubnetId string `json:"subnet_id"`

	// 本端子网列表
	LocalSubnetList []string `json:"local_subnet_list"`

	// 虚拟私有云区域项目ID
	RegionProjectId string `json:"region_project_id"`

	// 区域ID
	RegionId string `json:"region_id"`
}

func (o AddEcnWithVpcRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEcnWithVpcRequestBody struct{}"
	}

	return strings.Join([]string{"AddEcnWithVpcRequestBody", string(data)}, " ")
}
