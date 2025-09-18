package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStrategyResponse Response Object
type UpdateStrategyResponse struct {

	// **参数解释**： 是否调用成功。 **取值范围**： - true：调用成功。 - false：调用失败。
	Status *bool `json:"status,omitempty"`

	// **参数解释**： 策略ID，策略的唯一标识，通过[获取策略列表](ListStrategy.xml)接口获取，data.id即为策略ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	RuleSetId      *string `json:"rule_set_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStrategyResponse struct{}"
	}

	return strings.Join([]string{"UpdateStrategyResponse", string(data)}, " ")
}
