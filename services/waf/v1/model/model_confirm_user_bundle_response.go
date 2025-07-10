package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmUserBundleResponse Response Object
type ConfirmUserBundleResponse struct {

	// **参数解释：** 云模式套餐类型 **取值范围：** - -2：冻结 - -1：无 - 2：云模式 标准版（包周期） - 3：云模式 专业版（包周期） - 4：云模式 企业版（包周期） - 7：云模式 入门版（包周期） - 22：云模式（按需计费）
	Type *int32 `json:"type,omitempty"`

	// **参数解释：** 云模式套餐名称 **取值范围：** - None：无 - Basic：云模式 入门版（包周期） - Professional：云模式 标准版（包周期） - Enterprise：云模式 专业版（包周期） - Ultimate: 云模式 企业版（包周期） - cloud.waf.postpaid：云模式（按需计费）
	Name *string `json:"name,omitempty"`

	// **参数解释：** 云模式支持的域名配额相关信息 **取值范围：** 不涉及
	Host *interface{} `json:"host,omitempty"`

	// **参数解释：** 独享套餐类型 **取值范围：** - -2：冻结 - -1：无 - 12：独享模式 版本规格为WI-100 - 13：独享模式 版本规格为WI-500
	PremiumType *int32 `json:"premium_type,omitempty"`

	// **参数解释：** 独享模式套餐名称 **取值范围：** - None：无 - Instance.professional：独享模式 版本规格为WI-100 - Instance.enterprise：独享模式 版本规格为WI-500
	PremiumName *string `json:"premium_name,omitempty"`

	// **参数解释：** 独享支持的域名配额相关信息 **取值范围：** 不涉及
	PremiumHost *interface{} `json:"premium_host,omitempty"`

	// **参数解释：** 支持的策略相关信息 **取值范围：** 不涉及
	Options *interface{} `json:"options,omitempty"`

	// **参数解释：** 支持的规则配额相关信息 **取值范围：** 不涉及
	Rule *interface{} `json:"rule,omitempty"`

	// **参数解释：** 不同版本支持的规则信息 **取值范围：** 不涉及
	Upgrade *interface{} `json:"upgrade,omitempty"`

	// **参数解释：** 支持的特性 **取值范围：** 不涉及
	Feature        *interface{} `json:"feature,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ConfirmUserBundleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmUserBundleResponse struct{}"
	}

	return strings.Join([]string{"ConfirmUserBundleResponse", string(data)}, " ")
}
