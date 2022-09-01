package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgencyV2 struct {

	// 委托id
	Id *string `json:"id,omitempty" xml:"id"`

	// 委托名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 委托的期限
	Duration *string `json:"duration,omitempty" xml:"duration"`

	// 委托的账号名称
	TrustDomainName *string `json:"trust_domain_name,omitempty" xml:"trust_domain_name"`

	// 委托的账号id
	TrustDomainId *string `json:"trust_domain_id,omitempty" xml:"trust_domain_id"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 超期时间
	ExpireTime *string `json:"expire_time,omitempty" xml:"expire_time"`
}

func (o AgencyV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyV2 struct{}"
	}

	return strings.Join([]string{"AgencyV2", string(data)}, " ")
}
