package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type IdentityAssumerole struct {

	// 委托名。
	AgencyName string `json:"agency_name" xml:"agency_name"`

	// 委托方的账号ID。“domain_id”与“domain_name”至少填写一个。
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 委托方的账号名。“domain_id”与“domain_name”至少填写一个。
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// AK/SK和securitytoken的有效期，时间单位为秒。取值范围：15min ~ 24h ，默认为15min。
	DurationSeconds *int32 `json:"duration_seconds,omitempty" xml:"duration_seconds"`

	SessionUser *AssumeroleSessionuser `json:"session_user,omitempty" xml:"session_user"`
}

func (o IdentityAssumerole) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdentityAssumerole struct{}"
	}

	return strings.Join([]string{"IdentityAssumerole", string(data)}, " ")
}
