package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFirewallListResponse Response Object
type ListFirewallListResponse struct {

	// 是否支持eps
	UserSupportEps *bool `json:"user_support_eps,omitempty"`

	// 是否存在ndr
	HasNdr *bool `json:"has_ndr,omitempty"`

	// 是否支持按需购买
	IsSupportPostpaid *bool `json:"is_support_postpaid,omitempty"`

	// 是否支持基础版
	IsSupportBasicVersion *bool `json:"is_support_basic_version,omitempty"`

	// 是否支持购买专业版
	IsSupportBuyProfessional *bool `json:"is_support_buy_professional,omitempty"`

	Data           *HttpFirewallInstanceListResponseData `json:"data,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListFirewallListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirewallListResponse struct{}"
	}

	return strings.Join([]string{"ListFirewallListResponse", string(data)}, " ")
}
