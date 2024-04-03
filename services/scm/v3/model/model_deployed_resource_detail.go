package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeployedResourceDetail struct {

	// 证书已部署资源服务名称。 - WAF：证书关联Web应用防火墙的资源。 - CDN：证书关联内容分发网络的资源。 - ELB：证书关联弹性负载均衡（经典型）的资源。
	Service string `json:"service"`

	// 证书在当前服务已部署资源数量。
	ResourceNum int32 `json:"resource_num"`

	// 全局服务或Region级服务。
	ResourceLocation string `json:"resource_location"`

	// 局点资源列表，详情请参见RegionResourceDetail字段数据结构说明。
	RegionResources []RegionResourceDetail `json:"region_resources"`
}

func (o DeployedResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployedResourceDetail struct{}"
	}

	return strings.Join([]string{"DeployedResourceDetail", string(data)}, " ")
}
