package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgencyV2 struct {

	// 委托id
	Id *string `json:"id,omitempty"`

	// 委托名称
	Name *string `json:"name,omitempty"`

	// 委托的期限
	Duration *string `json:"duration,omitempty"`

	// 委托的账号名称
	TrustDomainName *string `json:"trust_domain_name,omitempty"`

	// 委托的账号id
	TrustDomainId *string `json:"trust_domain_id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 超期时间
	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o AgencyV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyV2 struct{}"
	}

	return strings.Join([]string{"AgencyV2", string(data)}, " ")
}
