package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePunishmentRuleResponse Response Object
type UpdatePunishmentRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 所属策略id
	Policyid *string `json:"policyid,omitempty"`

	// 拦截时间
	BlockTime *int32 `json:"block_time,omitempty"`

	// 攻击惩罚类别
	Category *string `json:"category,omitempty"`

	// 规则描述
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePunishmentRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePunishmentRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdatePunishmentRuleResponse", string(data)}, " ")
}
