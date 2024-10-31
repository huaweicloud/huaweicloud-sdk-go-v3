package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowHttpPunishmentRuleResponseBody struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 规则所在策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// 创建规则时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 拦截类型，可选值为：long_ip_block（长时间IP拦截）、long_cookie_block（长时间Cookie拦截）、long_params_block（长时间Params拦截）、short_ip_block（短时间IP拦截）、short_cookie_block（短时间Cookie拦截）、short_params_block（短时间Params拦截）
	Category *string `json:"category,omitempty"`

	// 拦截时长
	BlockTime *int32 `json:"block_time,omitempty"`
}

func (o ShowHttpPunishmentRuleResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpPunishmentRuleResponseBody struct{}"
	}

	return strings.Join([]string{"ShowHttpPunishmentRuleResponseBody", string(data)}, " ")
}
