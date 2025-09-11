package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SignSpecialAgreementResponse Response Object
type SignSpecialAgreementResponse struct {

	// 当前服务协议类型
	AgreementType *string `json:"agreement_type,omitempty"`

	// 当前服务协议版本
	CurrentVersion *string `json:"current_version,omitempty"`

	// 签署服务协议版本
	SignedVersion *string `json:"signed_version,omitempty"`

	// 协议签署时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	SignedTime *string `json:"signed_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SignSpecialAgreementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignSpecialAgreementResponse struct{}"
	}

	return strings.Join([]string{"SignSpecialAgreementResponse", string(data)}, " ")
}
