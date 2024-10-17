package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuditSqlResponseSql sql信息
type AuditSqlResponseSql struct {

	// ID
	Id string `json:"id"`

	// sql语句
	SqlStatement *string `json:"sql_statement,omitempty"`

	// 客户端IP
	ClientIp *string `json:"client_ip,omitempty"`

	// 客户端名称
	ClientName *string `json:"client_name,omitempty"`

	// 数据库IP
	DbIp *string `json:"db_ip,omitempty"`

	// 数据库用户名
	DbUser *string `json:"db_user,omitempty"`

	// 查询类型 LOGIN,CREATE_TABLE,CREATE_TABLESPACE,DROP_TABLE, DROP_TABLESPACE,DELETE,INSERT,INSERT_SELECT,SELECT,SELECT_FOR_UPDATE, UPDATE,CREATE_USER,DROP_USER,GRANT,OPERATE ALL
	QueryType *string `json:"query_type,omitempty"`

	// 操作对象列表
	OperatedObjInfo *[]AuditSqlResponseSqlOperatedObjInfo `json:"operated_obj_info,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 风险级别 - HIGH - MEDIUM - LOW - NO_RISK
	RiskLevel *string `json:"risk_level,omitempty"`

	// 审计开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 响应结果 - SUCCESS - FAILED
	SqlResponse string `json:"sql_response"`

	// 数据库实例
	DbInstance *string `json:"db_instance,omitempty"`
}

func (o AuditSqlResponseSql) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditSqlResponseSql struct{}"
	}

	return strings.Join([]string{"AuditSqlResponseSql", string(data)}, " ")
}
