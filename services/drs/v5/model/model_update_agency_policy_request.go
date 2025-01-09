package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAgencyPolicyRequest Request Object
type UpdateAgencyPolicyRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 委托名称。
	AgencyName string `json:"agency_name"`

	Body *UpdateAgencyPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateAgencyPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateAgencyPolicyRequest", string(data)}, " ")
}
