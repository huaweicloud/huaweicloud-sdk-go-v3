package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateHttpPunishmentRuleRequestBody struct {

	// 规则描述，最长512字符
	Description *string `json:"description,omitempty"`

	// 拦截时间，如果选择前缀为long的攻击惩罚类别，则block_time时长范围设置为301-1800;选择前缀为short的攻击惩罚类别，则block_time时长范围为1-300之间
	BlockTime int32 `json:"block_time"`
}

func (o UpdateHttpPunishmentRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpPunishmentRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHttpPunishmentRuleRequestBody", string(data)}, " ")
}
