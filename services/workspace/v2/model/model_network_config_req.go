package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkConfigReq 网络配置。
type NetworkConfigReq struct {

	// VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 指定业务子网的网络id。
	SubnetIds *[]string `json:"subnet_ids,omitempty"`
}

func (o NetworkConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkConfigReq struct{}"
	}

	return strings.Join([]string{"NetworkConfigReq", string(data)}, " ")
}
