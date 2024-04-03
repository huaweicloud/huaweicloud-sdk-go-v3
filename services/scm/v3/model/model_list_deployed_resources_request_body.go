package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListDeployedResourcesRequestBody struct {

	// 证书ID列表。
	CertificateIds []string `json:"certificate_ids"`

	// 服务名称列表。 - WAF：查询证书关联Web应用防火墙的资源。 - CDN：查询证书关联内容分发网络的资源。 - ELB：查询证书关联弹性负载均衡（经典型）的资源。 - OSB：查询证书关联对象存储服务的资源。 - ALL：查询证书以上四种服务的资源。
	ServiceNames []string `json:"service_names"`
}

func (o ListDeployedResourcesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeployedResourcesRequestBody struct{}"
	}

	return strings.Join([]string{"ListDeployedResourcesRequestBody", string(data)}, " ")
}
