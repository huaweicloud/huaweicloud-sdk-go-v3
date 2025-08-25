package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociationInfo struct {

	// 绑定关系ID
	Id string `json:"id"`

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 集群名称
	ClusterName string `json:"cluster_name"`

	// 应用ID
	AppId string `json:"app_id"`

	// 应用名称
	AppName string `json:"app_name"`

	// 应用所属的VPC名称
	VpcName string `json:"vpc_name"`

	// 应用所属的VPC的子网
	SubnetName string `json:"subnet_name"`

	// 集群所属的服务
	ClusterServerType string `json:"cluster_server_type"`

	// 访问地址
	VpcepAddress string `json:"vpcep_address"`

	// 绑定关系的最新更改时间，UNIX时间戳，单位毫秒
	UpdateTime int64 `json:"update_time"`

	// 绑定关系的创建时间，UNIX时间戳，单位毫秒
	CreateTime int64 `json:"create_time"`
}

func (o AssociationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociationInfo struct{}"
	}

	return strings.Join([]string{"AssociationInfo", string(data)}, " ")
}
