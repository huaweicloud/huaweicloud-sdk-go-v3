package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCustomRuleConfigResponse struct {

	// **参数解释**： 规则ID **取值范围**： 字符长度1-36位
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释**: 防护主机数量。 **取值范围**: 最小值1，最大值2000000
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**： 规则名称 **取值范围**： 字符长度1-64位
	RuleName *string `json:"rule_name,omitempty"`

	// **参数解释**： 规则状态 **取值范围**: - 0：停用 - 1：启用
	RuleStatus *int32 `json:"rule_status,omitempty"`

	// **参数解释**： 规则类型 **取值范围**： - black_hash：黑hash
	RuleType *string `json:"rule_type,omitempty"`

	// **参数解释**： 是否自动阻断告警 **取值范围**： - 0：不自动阻断告警 - 1：自动阻断告警
	AutoBlock *int32 `json:"auto_block,omitempty"`

	// **参数解释**： hash类型 **取值范围**： - SHA-256：sha256sum - MD5：md5sum - SHA-1：sha1sum
	HashType *string `json:"hash_type,omitempty"`

	// **参数解释**: 是否选择所有主机 **取值范围**: - true：是 - false：否
	IsAllHost *bool `json:"is_all_host,omitempty"`

	// **参数解释**： 创建时间 **取值范围**： 最小值0，最大值9223372036854775807
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 创建时间 **取值范围**： 最小值0，最大值9223372036854775807
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o ListCustomRuleConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomRuleConfigResponse struct{}"
	}

	return strings.Join([]string{"ListCustomRuleConfigResponse", string(data)}, " ")
}
