package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWafPolicyResponse Response Object
type ShowWafPolicyResponse struct {

	// 域名(包含端口)
	DomainName *string `json:"domain_name,omitempty"`

	// 0-中国大陆，1-中国大陆外
	OverseasType *int32 `json:"overseas_type,omitempty"`

	Options *WafPolicyOptions `json:"options,omitempty"`

	// 智能CC防护等级：[0-宽松,1- 正常, 2- 严格]
	Level *int32 `json:"level,omitempty"`

	// 智能CC模式：0-预警，1-防护
	Mode           *int32 `json:"mode,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowWafPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWafPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowWafPolicyResponse", string(data)}, " ")
}
