package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAntileakageRuleResponse Response Object
type UpdateAntileakageRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 规则应用的url
	Url *string `json:"url,omitempty"`

	// 类别（响应码：code，敏感信息：sensitive）
	Category *string `json:"category,omitempty"`

	// 规则内容
	Contents *[]string `json:"contents,omitempty"`

	// **参数解释：** 规则状态标识，用于指定规则的启用或关闭状态 **约束限制：** 不涉及 **取值范围：**  - 0：关闭  - 1：开启 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// 规则描述
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAntileakageRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAntileakageRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateAntileakageRuleResponse", string(data)}, " ")
}
