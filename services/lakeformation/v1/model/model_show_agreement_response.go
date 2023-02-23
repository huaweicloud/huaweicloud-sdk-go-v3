package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAgreementResponse struct {

	// 租户协议列表
	Agreements *[]TenantAgreement `json:"agreements,omitempty"`

	// 是否绑定了委托
	IsAgency *bool `json:"is_agency,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAgreementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgreementResponse struct{}"
	}

	return strings.Join([]string{"ShowAgreementResponse", string(data)}, " ")
}
