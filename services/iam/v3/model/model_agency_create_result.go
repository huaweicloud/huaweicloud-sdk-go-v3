package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyCreateResult
type AgencyCreateResult struct {

	// 委托创建时间。
	CreateTime string `json:"create_time"`

	// 委托描述信息。
	Description string `json:"description"`

	// 委托方账号ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 委托的期限。取值为\"FOREVER\"或“null”表示委托的期限为永久，取值为24表示委托的期限为一天，或为24小时。
	Duration string `json:"duration"`

	// 委托过期时间。“null”表示不过期。
	ExpireTime string `json:"expire_time"`

	// 委托ID。
	Id string `json:"id"`

	// 委托名。
	Name string `json:"name"`

	// 被委托方账号ID。
	TrustDomainId *string `json:"trust_domain_id,omitempty"`
}

func (o AgencyCreateResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyCreateResult struct{}"
	}

	return strings.Join([]string{"AgencyCreateResult", string(data)}, " ")
}
