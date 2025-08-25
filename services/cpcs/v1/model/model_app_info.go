package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppInfo struct {

	// 应用ID
	AppId string `json:"app_id"`

	// 应用名称
	AppName string `json:"app_name"`

	// 应用所属的VPC的ID
	VpcId string `json:"vpc_id"`

	// 应用所属的VPC的名称
	VpcName string `json:"vpc_name"`

	// 应用所属的VPC的子网iID
	SubnetId string `json:"subnet_id"`

	// 应用所属的VPC的子网
	SubnetName string `json:"subnet_name"`

	// 账号ID
	DomainId string `json:"domain_id"`

	// 应用描述
	Description string `json:"description"`

	// 应用的创建时间，UNIX时间戳，单位毫秒
	CreateTime int64 `json:"create_time"`
}

func (o AppInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppInfo struct{}"
	}

	return strings.Join([]string{"AppInfo", string(data)}, " ")
}
