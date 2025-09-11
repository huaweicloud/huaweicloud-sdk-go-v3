package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleAddRiskRuleRequest struct {

	// 操作类型，多个用英文','分隔
	Action *string `json:"action,omitempty"`

	// 影响行数
	AffectRows int32 `json:"affect_rows"`

	// 影响行数操作符 - GREATER: 大于 - EQUAL: 等于 - LESS: 小于 - GREATER_EQUAL: 大于等于 - LESS_EQUAL: 小于等于 - NO_MATCH: 不等于
	AffectSymbol string `json:"affect_symbol"`

	// 客户端IP，多个用英文','分隔
	ClientIps *string `json:"client_ips,omitempty"`

	// 数据库ID，多个用英文','分隔
	DbIds *string `json:"db_ids,omitempty"`

	// 例外IP，IP内不做匹配，多个用英文','分隔
	ExceptionIps *string `json:"exception_ips,omitempty"`

	// 执行时长操作符 - GREATER: 大于 - EQUAL: 等于 - LESS: 小于 - GREATER_EQUAL: 大于等于 - LESS_EQUAL: 小于等于 - NO_MATCH: 不等于
	ExecutionSymbol string `json:"execution_symbol"`

	// 执行时长
	ExecutionTime int64 `json:"execution_time"`

	// 是否忽略大小写
	IgnoreCase *bool `json:"ignore_case,omitempty"`

	// 风险等级   - LOW：低  - MEDIUM：中  - HIGH：高  - NO_RISK：无
	RiskLevel string `json:"risk_level"`

	// 规则名称
	RuleName string `json:"rule_name"`

	// 操作对象
	Schemas *[]SchemaBean `json:"schemas,omitempty"`

	// 状态  - OFF：已关闭  - ON：已开启
	Status string `json:"status"`
}

func (o RuleAddRiskRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleAddRiskRuleRequest struct{}"
	}

	return strings.Join([]string{"RuleAddRiskRuleRequest", string(data)}, " ")
}
