package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群负载均衡信息。
type CreateClusterLoadBalance struct {
	// 是否开启内网域名。

	EndpointWithDnsName bool `json:"endpointWithDnsName"`
	// 访问控制。

	VpcPermisssions *[]string `json:"vpcPermisssions,omitempty"`
}

func (o CreateClusterLoadBalance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterLoadBalance struct{}"
	}

	return strings.Join([]string{"CreateClusterLoadBalance", string(data)}, " ")
}
