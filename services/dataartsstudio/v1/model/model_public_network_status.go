package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicNetworkStatus 集群认证信息
type PublicNetworkStatus struct {

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 是否开启了公网访问,true:开启，false:未开启
	Status *bool `json:"status,omitempty"`

	// 公网ip地址
	Ip *string `json:"ip,omitempty"`
}

func (o PublicNetworkStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicNetworkStatus struct{}"
	}

	return strings.Join([]string{"PublicNetworkStatus", string(data)}, " ")
}
