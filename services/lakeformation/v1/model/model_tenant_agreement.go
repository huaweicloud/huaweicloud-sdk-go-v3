package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TenantAgreement 租户协议
type TenantAgreement struct {

	// 协议名称
	AgreementName *string `json:"agreement_name,omitempty"`

	// 协议版本号
	AgreementVersion *string `json:"agreement_version,omitempty"`
}

func (o TenantAgreement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantAgreement struct{}"
	}

	return strings.Join([]string{"TenantAgreement", string(data)}, " ")
}
