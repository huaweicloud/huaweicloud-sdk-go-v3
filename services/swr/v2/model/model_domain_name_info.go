package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainNameInfo struct {

	// 域名ID
	Uid *string `json:"uid,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 域名类型(default/custom)
	Type *string `json:"type,omitempty"`

	// 域名，只能由字母、数字、中划线、星号组成， 星号只能在开头，中划线不能在开头或末 尾，至少包含两个字符串，单个字符串不 超过63个字符，字符串间以点分割，且总 长度不超过100个字符。例如： example.com 或 *example.com。
	DomainName *string `json:"domain_name,omitempty"`

	// SCM服务的证书ID
	CertificateId *string `json:"certificate_id,omitempty"`

	// 增加域名的时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新域名的时间
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o DomainNameInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainNameInfo struct{}"
	}

	return strings.Join([]string{"DomainNameInfo", string(data)}, " ")
}
