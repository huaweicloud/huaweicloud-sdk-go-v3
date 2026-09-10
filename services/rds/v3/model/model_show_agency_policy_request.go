package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgencyPolicyRequest Request Object
type ShowAgencyPolicyRequest struct {

	// 委托名称。目前仅支持RDSAccessProjectResource。
	AgencyName string `json:"agency_name"`

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowAgencyPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowAgencyPolicyRequest", string(data)}, " ")
}
