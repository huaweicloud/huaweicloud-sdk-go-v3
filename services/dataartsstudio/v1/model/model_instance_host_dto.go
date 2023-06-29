package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceHostDto struct {

	// 集群id
	InstanceId *string `json:"instance_id,omitempty"`

	// 集群名
	InstanceName *string `json:"instance_name,omitempty"`

	// 内网地址
	IntranetHost *string `json:"intranet_host,omitempty"`

	// 外网地址
	ExternalHost *string `json:"external_host,omitempty"`

	// 网关域名
	Domains *[]string `json:"domains,omitempty"`
}

func (o InstanceHostDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceHostDto struct{}"
	}

	return strings.Join([]string{"InstanceHostDto", string(data)}, " ")
}
