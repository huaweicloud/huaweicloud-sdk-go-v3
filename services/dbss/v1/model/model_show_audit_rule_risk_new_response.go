package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditRuleRiskNewResponse Response Object
type ShowAuditRuleRiskNewResponse struct {

	// 风险规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 风险名称
	RuleName *string `json:"rule_name,omitempty"`

	// 风险规则状态 - OFF：禁用 - ON：启用
	Status *string `json:"status,omitempty"`

	// 操作集合, 中间逗号分隔。 LOGIN,CREATE_TABLE,CREATE_TABLESPACE,DROP_TABLE, DROP_TABLESPACE,DELETE,INSERT,INSERT_SELECT,SELECT,SELECT_FOR_UPDATE, UPDATE,CREATE_USER,DROP_USER,GRANT,OPERATE ALL
	Action *string `json:"action,omitempty"`

	// Schema列表
	Schemas *[]RuleRiskInfoBeanSchemas `json:"schemas,omitempty"`

	// 风险规则优先级。数字越小优先级越高。
	Rank *int32 `json:"rank,omitempty"`

	// 是否忽略大小写
	IgnoreCase *bool `json:"ignore_case,omitempty"`

	// 风险级别 - LOW：低 - MEDIUM：中 - HIGH：高 - NO_RISK：无风险
	RiskLevel *string `json:"risk_level,omitempty"`

	// 数据库id，中间逗号分隔（单个id 小于256位）
	DbIds *string `json:"db_ids,omitempty"`

	// 执行时长对执行时长阈值的关系 - GREATER：大于 - EQUAL：等于 - LESS：小于 - GREATER_EQUAL：大于等于 - LESS_EQUAL：小于等于 - NO_MATCH：不匹配
	ExecutionSymbol *string `json:"execution_symbol,omitempty"`

	// 设定的执行时长阈值
	ExecutionTime *int32 `json:"execution_time,omitempty"`

	// 影响行数对行数阈值的关系：  - GREATER：大于 - EQUAL：等于 - LESS：小于 - GREATER_EQUAL：大于等于 - LESS_EQUAL：小于等于 - NO_MATCH：不匹配
	AffectSymbol *string `json:"affect_symbol,omitempty"`

	// 设定的影响行数阈值
	AffectRows *int32 `json:"affect_rows,omitempty"`

	// 客户端IP段: IP-IP格式，或IP/XX 格式。 各个IP段使用逗号连接
	ClientIps      *string `json:"client_ips,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAuditRuleRiskNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditRuleRiskNewResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditRuleRiskNewResponse", string(data)}, " ")
}
