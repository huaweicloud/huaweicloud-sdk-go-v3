package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyUpdateResult
type AgencyUpdateResult struct {

	// 委托创建时间。
	CreateTime string `json:"create_time"`

	// 委托描述信息。
	Description string `json:"description"`

	// 委托方账号ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 委托过期时间。“null”表示不过期。
	ExpireTime string `json:"expire_time"`

	// 委托ID。
	Id string `json:"id"`

	// 委托名。
	Name string `json:"name"`

	// 被委托方账号ID。
	TrustDomainId *string `json:"trust_domain_id,omitempty"`
}

func (o AgencyUpdateResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyUpdateResult struct{}"
	}

	return strings.Join([]string{"AgencyUpdateResult", string(data)}, " ")
}
