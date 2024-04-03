package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeployCertificateRequestBody struct {

	// 部署的资源所在的项目名称，若在主项目下，则该值为region id。
	ProjectName *string `json:"project_name,omitempty"`

	// 证书推送的目标服务，当前仅支持：CDN、WAF、ELB。
	ServiceName string `json:"service_name"`

	// 所要部署的资源列表。
	Resources []DeployedResource `json:"resources"`
}

func (o DeployCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"DeployCertificateRequestBody", string(data)}, " ")
}
