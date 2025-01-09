package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSubnetIdsRequestBody struct {

	// 业务子网id
	SubnetIds *[]string `json:"subnet_ids,omitempty"`

	// VPC配置信息列表。
	VpcConfigInfos *[]VpcConfigInfo `json:"vpc_config_infos,omitempty"`
}

func (o UpdateSubnetIdsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetIdsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSubnetIdsRequestBody", string(data)}, " ")
}
