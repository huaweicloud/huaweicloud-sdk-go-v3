package model

import (
	"encoding/json"

	"strings"
)

type ShowQuotaResponseBodyQuotas struct {
	// 配额上限

	QuotaLimit *int32 `json:"quota_limit,omitempty"`
	// 配额类型

	Type *string `json:"type,omitempty"`
	// 已使用配额数

	Used *int32 `json:"used,omitempty"`
	// 域名所属用户的domain_id。

	UserDomainId *string `json:"user_domain_id,omitempty"`
}

func (o ShowQuotaResponseBodyQuotas) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQuotaResponseBodyQuotas struct{}"
	}

	return strings.Join([]string{"ShowQuotaResponseBodyQuotas", string(data)}, " ")
}
