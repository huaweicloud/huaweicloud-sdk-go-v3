package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIgnoreRuleResponse Response Object
type DeleteIgnoreRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 创建规则的时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// **参数解释：** 规则状态标识，用于指定规则的启用或关闭状态 **约束限制：** 不涉及 **取值范围：**  - 0：关闭  - 1：开启 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// 误报规则屏蔽路径，仅在mode为0的状态下有该字段
	Url *string `json:"url,omitempty"`

	// 被屏蔽检测的规则类型或规则ID
	Rule *string `json:"rule,omitempty"`

	// 版本号，0代表v1旧版本，1代表v2新版本；mode为0时，不存在conditions字段，存在url和url_logic字段；mode为1时，不存在url和url_logic字段，存在conditions字段
	Mode *int32 `json:"mode,omitempty"`

	// url匹配逻辑
	UrlLogic *string `json:"url_logic,omitempty"`

	// 条件
	Conditions *[]Condition `json:"conditions,omitempty"`

	Advanced *IgnoreAdvanced `json:"advanced,omitempty"`

	// 防护域名或防护网站
	Domain         *[]string `json:"domain,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteIgnoreRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIgnoreRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteIgnoreRuleResponse", string(data)}, " ")
}
