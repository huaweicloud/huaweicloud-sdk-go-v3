package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditRuleRiskResponse Response Object
type ShowAuditRuleRiskResponse struct {

	// 风险规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 风险名称
	RuleName *string `json:"rule_name,omitempty"`

	// 风险规则状态 - OFF - ON
	Status *string `json:"status,omitempty"`

	// 操作集合, 中间逗号分隔。 LOGIN,CREATE_TABLE,CREATE_TABLESPACE,DROP_TABLE, DROP_TABLESPACE,DELETE,INSERT,INSERT_SELECT,SELECT,SELECT_FOR_UPDATE, UPDATE,CREATE_USER,DROP_USER,GRANT,OPERATE ALL
	Action *string `json:"action,omitempty"`

	// Schema列表
	Schemas *[]RuleRiskInfoBeanSchemas `json:"schemas,omitempty"`

	// 风险规则优先级。数字越小优先级越高。
	Rank *int32 `json:"rank,omitempty"`

	// 是否忽略大小写
	IgnoreCase *bool `json:"ignore_case,omitempty"`

	// 风险级别 - LOW - MEDIUM - HIGH - NO_RISK
	RiskLevel *string `json:"risk_level,omitempty"`

	// 数据库id，中间逗号分隔（单个id 小于256位）
	DbIds *string `json:"db_ids,omitempty"`

	// 执行时长对执行时长阈值的关系 - GREATER - EQUAL - LESS - GREATER_EQUAL - LESS_EQUAL - NO_MATCH
	ExecutionSymbol *string `json:"execution_symbol,omitempty"`

	// 设定的执行时长阈值
	ExecutionTime *int32 `json:"execution_time,omitempty"`

	// 影响行数对行数阈值的关系：  - GREATER - EQUAL - LESS - GREATER_EQUAL - LESS_EQUAL - NO_MATCH
	AffectSymbol *string `json:"affect_symbol,omitempty"`

	// 设定的影响行数阈值
	AffectRows *int32 `json:"affect_rows,omitempty"`

	// 客户端IP段: IP-IP格式，或IP/XX 格式。 各个IP段使用逗号连接
	ClientIps      *string `json:"client_ips,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAuditRuleRiskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditRuleRiskResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditRuleRiskResponse", string(data)}, " ")
}
