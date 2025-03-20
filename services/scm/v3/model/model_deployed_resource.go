package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeployedResource struct {

	// 资源Id，部署WAF与ELB时，必传此字段。
	Id *string `json:"id,omitempty"`

	// 资源类型，当前仅部署WAF资源时需传入，即独享模式（premium）与云模式（cloud）。
	Type *string `json:"type,omitempty"`

	// 需部署的域名，当前仅部署CDN时需传入，即需加速的域名，域名与证书必须可匹配。
	DomainName *string `json:"domain_name,omitempty"`

	// 需部署的资源所属的企业项目ID，当前仅部署WAF资源时，需传入。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o DeployedResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployedResource struct{}"
	}

	return strings.Join([]string{"DeployedResource", string(data)}, " ")
}
