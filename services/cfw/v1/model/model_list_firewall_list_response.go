package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFirewallListResponse Response Object
type ListFirewallListResponse struct {

	// **参数解释**： 是否支持企业项目 **取值范围**： - true：是 - false：否
	UserSupportEps *bool `json:"user_support_eps,omitempty"`

	// **参数解释**： 是否存在NDR，NDR为原旁路版防火墙，现已停止售卖。 **取值范围**： - true：是 - false：不是
	HasNdr *bool `json:"has_ndr,omitempty"`

	// **参数解释**： 是否支持按需购买 **取值范围**： - true：是 - false：不是
	IsSupportPostpaid *bool `json:"is_support_postpaid,omitempty"`

	// **参数解释**： 是否支持基础版 **取值范围**： - true：是 - false：不是
	IsSupportBasicVersion *bool `json:"is_support_basic_version,omitempty"`

	// **参数解释**： 是否支持购买专业版 **取值范围**： - true：是 - false：不是
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
