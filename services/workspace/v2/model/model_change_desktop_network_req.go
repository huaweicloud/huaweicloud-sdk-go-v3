package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDesktopNetworkReq 切换桌面网络请求体
type ChangeDesktopNetworkReq struct {

	// 待切换VPC的ID
	VpcId string `json:"vpc_id"`

	// 待切换子网的ID
	SubnetId string `json:"subnet_id"`

	// 指定私有IP地址
	PrivateIp *string `json:"private_ip,omitempty"`

	// 安全组ID列表
	SecurityGroupIds []string `json:"security_group_ids"`
}

func (o ChangeDesktopNetworkReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDesktopNetworkReq struct{}"
	}

	return strings.Join([]string{"ChangeDesktopNetworkReq", string(data)}, " ")
}
