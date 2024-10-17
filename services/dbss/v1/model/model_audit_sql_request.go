package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AuditSqlRequest struct {
	Time *AuditSqlRequestTime `json:"time"`

	// 风险级别 - HIGH - MEDIUM - LOW - NO_RISK
	RiskLevels *AuditSqlRequestRiskLevels `json:"risk_levels,omitempty"`

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
	Page *int32 `json:"page,omitempty"`

	// 条数
	Size *int32 `json:"size,omitempty"`

	// 时间顺序 - DESC - ASC
	TimeOrder *string `json:"time_order,omitempty"`
}

func (o AuditSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditSqlRequest struct{}"
	}

	return strings.Join([]string{"AuditSqlRequest", string(data)}, " ")
}

type AuditSqlRequestRiskLevels struct {
	value string
}

type AuditSqlRequestRiskLevelsEnum struct {
	HIGH    AuditSqlRequestRiskLevels
	MEDIUM  AuditSqlRequestRiskLevels
	LOW     AuditSqlRequestRiskLevels
	NO_RISK AuditSqlRequestRiskLevels
}

func GetAuditSqlRequestRiskLevelsEnum() AuditSqlRequestRiskLevelsEnum {
	return AuditSqlRequestRiskLevelsEnum{
		HIGH: AuditSqlRequestRiskLevels{
			value: "HIGH",
		},
		MEDIUM: AuditSqlRequestRiskLevels{
			value: "MEDIUM",
		},
		LOW: AuditSqlRequestRiskLevels{
			value: "LOW",
		},
		NO_RISK: AuditSqlRequestRiskLevels{
			value: "NO_RISK",
		},
	}
}

func (c AuditSqlRequestRiskLevels) Value() string {
	return c.value
}

func (c AuditSqlRequestRiskLevels) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuditSqlRequestRiskLevels) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
