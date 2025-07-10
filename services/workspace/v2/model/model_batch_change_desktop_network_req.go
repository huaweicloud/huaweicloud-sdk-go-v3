package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchChangeDesktopNetworkReq struct {

	// 桌面id列表，最小为1，最大为100。
	DesktopIds []string `json:"desktop_ids"`

	// 待切换VPC的ID。
	VpcId string `json:"vpc_id"`

	// 待切换子网的ID。
	SubnetId string `json:"subnet_id"`

	// 安全组ID列表。
	SecurityGroupIds *[]string `json:"security_group_ids,omitempty"`
}

func (o BatchChangeDesktopNetworkReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeDesktopNetworkReq struct{}"
	}

	return strings.Join([]string{"BatchChangeDesktopNetworkReq", string(data)}, " ")
}
