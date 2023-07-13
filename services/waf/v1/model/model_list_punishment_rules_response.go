package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPunishmentRulesResponse Response Object
type ListPunishmentRulesResponse struct {

	// 攻击惩罚规则数量
	Total *int32 `json:"total,omitempty"`

	// 攻击惩罚规则列表
	Items          *[]PunishmentInfo `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListPunishmentRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPunishmentRulesResponse struct{}"
	}

	return strings.Join([]string{"ListPunishmentRulesResponse", string(data)}, " ")
}
