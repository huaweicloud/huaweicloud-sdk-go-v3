package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateHttpPunishmentRuleRequestBody struct {

	// 规则描述，最长512字符
	Description *string `json:"description,omitempty"`

	// 拦截类型，可选值为：long_ip_block（长时间IP拦截）、long_cookie_block（长时间Cookie拦截）、long_params_block（长时间Params拦截）、short_ip_block（短时间IP拦截）、short_cookie_block（短时间Cookie拦截）、short_params_block（短时间Params拦截）
	Category string `json:"category"`

	// 拦截时间，如果选择前缀为long的攻击惩罚类别，则block_time时长范围设置为301-1800;选择前缀为short的攻击惩罚类别，则block_time时长范围为0-300之间
	BlockTime int32 `json:"block_time"`
}

func (o CreateHttpPunishmentRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpPunishmentRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHttpPunishmentRuleRequestBody", string(data)}, " ")
}
