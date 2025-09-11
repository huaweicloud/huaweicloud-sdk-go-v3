package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditSqlRequest struct {
	Time *AuditSqlRequestTime `json:"time"`

	// 风险级别 - HIGH：高 - MEDIUM：中 - LOW：低 - NO_RISK：无
	RiskLevels *string `json:"risk_levels,omitempty"`

	// 客户端IP
	ClientIp *string `json:"client_ip,omitempty"`

	// 客户端名称
	ClientName *string `json:"client_name,omitempty"`

	// 数据库IP
	DbIp *string `json:"db_ip,omitempty"`

	// 数据库用户
	DbUser *string `json:"db_user,omitempty"`

	// 查询类型 LOGIN,CREATE_TABLE,CREATE_TABLESPACE,DROP_TABLE, DROP_TABLESPACE,DELETE,INSERT,INSERT_SELECT,SELECT,SELECT_FOR_UPDATE, UPDATE,CREATE_USER,DROP_USER,GRANT,OPERATE ALL
	QueryType *string `json:"query_type,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// sql语句
	SqlStatement *string `json:"sql_statement,omitempty"`

	// 响应结果 - SUCCESS - FAILED
	SqlResponse *string `json:"sql_response,omitempty"`

	// 页码
	Page int32 `json:"page"`

	// 条数
	Size int32 `json:"size"`

	// 时间顺序 - DESC - ASC
	TimeOrder string `json:"time_order"`
}

func (o AuditSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditSqlRequest struct{}"
	}

	return strings.Join([]string{"AuditSqlRequest", string(data)}, " ")
}
