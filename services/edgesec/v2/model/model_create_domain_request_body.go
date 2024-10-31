package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainRequestBody 创建防护域名的请求
type CreateDomainRequestBody struct {

	// 防护域名（可带端口）
	DomainName string `json:"domain_name"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id，默认为0
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 防护域名关联的策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// 域名描述
	Description *string `json:"description,omitempty"`
}

func (o CreateDomainRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDomainRequestBody", string(data)}, " ")
}
