package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlDetailBean SQL语句对象
type SqlDetailBean struct {

	// 审计结果
	AuditResult *string `json:"audit_result,omitempty"`

	// 客户端端口
	ClientPort *int32 `json:"client_port,omitempty"`

	// 客户端IP
	ClientIp *string `json:"client_ip,omitempty"`

	// 客户端MAC地址
	ClientMac *string `json:"client_mac,omitempty"`

	// 客户端名称
	ClientName *string `json:"client_name,omitempty"`

	// 客户端系统主机名称
	ClientOsName *string `json:"client_os_name,omitempty"`

	// 客户端操作系统用户名
	ClientOsUser *string `json:"client_os_user,omitempty"`

	// 客户端端口字符
	ClientPortStr *string `json:"client_port_str,omitempty"`

	// 客户端连接工具
	ClientTool *string `json:"client_tool,omitempty"`

	// 数据库端口
	DbPort *int32 `json:"db_port,omitempty"`

	// 数据库实例
	DbInstance *string `json:"db_instance,omitempty"`

	// 数据库IP
	DbIp *string `json:"db_ip,omitempty"`

	// 数据库MAC地址
	DbMac *string `json:"db_mac,omitempty"`

	// 数据库端口字符
	DbPortStr *string `json:"db_port_str,omitempty"`

	// 数据库服务名称
	DbServiceName *string `json:"db_service_name,omitempty"`

	// 数据库连接ID
	DbSessionId *string `json:"db_session_id,omitempty"`

	// 数据库类型
	DbType *string `json:"db_type,omitempty"`

	// 数据库用户
	DbUser *string `json:"db_user,omitempty"`

	// 请求结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 执行时长
	ExecuteTime *int64 `json:"execute_time,omitempty"`

	// 操作对象类型信息
	Field *string `json:"field,omitempty"`

	// 记录ID
	Id *string `json:"id,omitempty"`

	// 风险等级   - LOW：低  - MEDIUM：中  - HIGH：高  - NO_RISK：无
	Level *string `json:"level,omitempty"`

	// 影响行数
	Lines *string `json:"lines,omitempty"`

	// 登入登出结果
	LogResult *string `json:"log_result,omitempty"`

	// 操作对象
	Object *string `json:"object,omitempty"`

	// 操作对象类型
	ObjectType *string `json:"object_type,omitempty"`

	// 查询ID
	QueryId *string `json:"query_id,omitempty"`

	// 操作类型
	QueryType *string `json:"query_type,omitempty"`

	// 数据长度
	ResponseLength *int32 `json:"response_length,omitempty"`

	// 响应状态
	ResponseStatus *string `json:"response_status,omitempty"`

	// 规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 规则类型
	RuleType *string `json:"rule_type,omitempty"`

	// 数据库SCHEMA
	Schema *string `json:"schema,omitempty"`

	// 审计范围ID
	ScopeId *string `json:"scope_id,omitempty"`

	// 审计范围名称
	ScopeName *string `json:"scope_name,omitempty"`

	// SQL返回结果
	SqlResponse *string `json:"sql_response,omitempty"`

	// SQL处理结果
	SqlResult *string `json:"sql_result,omitempty"`

	// SQL语句内容
	SqlStatement *string `json:"sql_statement,omitempty"`

	// 请求开始时间
	StartTime *string `json:"start_time,omitempty"`

	// TCP连接ID
	TcpSessionId *string `json:"tcp_session_id,omitempty"`
}

func (o SqlDetailBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlDetailBean struct{}"
	}

	return strings.Join([]string{"SqlDetailBean", string(data)}, " ")
}
