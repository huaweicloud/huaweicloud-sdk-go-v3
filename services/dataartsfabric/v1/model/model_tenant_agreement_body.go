package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TenantAgreementBody 租户注册协议请求体
type TenantAgreementBody struct {

	// 租户协议列表
	Agreements *[]TenantAgreement `json:"agreements,omitempty"`
}

func (o TenantAgreementBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantAgreementBody struct{}"
	}

	return strings.Join([]string{"TenantAgreementBody", string(data)}, " ")
}
