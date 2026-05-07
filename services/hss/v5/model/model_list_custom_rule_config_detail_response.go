package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomRuleConfigDetailResponse Response Object
type ListCustomRuleConfigDetailResponse struct {

	// **参数解释**： 规则ID **取值范围**： 字符长度1-36位
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释**： hash类型 **取值范围**： - SHA-256：sha256sum - MD5：md5sum - SHA-1：sha1sum
	HashType *string `json:"hash_type,omitempty"`

	// **参数解释**： 是否自动阻断告警 **取值范围**： - 0：不自动阻断告警 - 1：自动阻断告警
	AutoBlock *int32 `json:"auto_block,omitempty"`

	// **参数解释**: 是否选择所有主机 **取值范围**: - true：是 - false：否
	IsAllHost *bool `json:"is_all_host,omitempty"`

	// **参数解释**： 规则类型 **取值范围**： - black_hash：黑hash
	RuleType *string `json:"rule_type,omitempty"`

	// **参数解释**： 规则集列表 **取值范围**: 1-1000个规则值
	RuleValues *[]string `json:"rule_values,omitempty"`

	// **参数解释**: agent列表 **取值范围**: 字符长度1-64位
	AgentIds       *[]string `json:"agent_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListCustomRuleConfigDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomRuleConfigDetailResponse struct{}"
	}

	return strings.Join([]string{"ListCustomRuleConfigDetailResponse", string(data)}, " ")
}
