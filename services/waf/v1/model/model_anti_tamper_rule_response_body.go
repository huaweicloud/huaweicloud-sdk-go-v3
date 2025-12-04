package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AntiTamperRuleResponseBody struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 该规则所属防护策略的id
	Policyid *string `json:"policyid,omitempty"`

	// 创建规则的时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 该规则备注
	Description *string `json:"description,omitempty"`

	// **参数解释：** 规则状态标识，用于指定规则的启用或关闭状态 **约束限制：** 不涉及 **取值范围：**  - 0：关闭  - 1：开启 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// 防篡改的域名
	Hostname *string `json:"hostname,omitempty"`

	// 防篡改的url
	Url *string `json:"url,omitempty"`
}

func (o AntiTamperRuleResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiTamperRuleResponseBody struct{}"
	}

	return strings.Join([]string{"AntiTamperRuleResponseBody", string(data)}, " ")
}
