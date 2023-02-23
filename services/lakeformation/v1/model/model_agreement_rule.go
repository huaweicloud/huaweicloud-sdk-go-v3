package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 协议规则信息
type AgreementRule struct {

	// 协议名称
	AgreementName *string `json:"agreement_name,omitempty"`

	// 协议展示名称
	AgreementDisplayName *string `json:"agreement_display_name,omitempty"`

	// 协议版本
	AgreementVersion *string `json:"agreement_version,omitempty"`

	// 协议链接
	AgreementUrl *string `json:"agreement_url,omitempty"`
}

func (o AgreementRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgreementRule struct{}"
	}

	return strings.Join([]string{"AgreementRule", string(data)}, " ")
}
