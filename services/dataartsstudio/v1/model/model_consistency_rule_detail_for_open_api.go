package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsistencyRuleDetailForOpenApi struct {

	// 子规则ID
	Id *int64 `json:"id,omitempty"`

	// 子规则名称
	SubRuleName *string `json:"sub_rule_name,omitempty"`

	// SingleDatabase:库级规则，SingleTable:表级规则,SingleColumn:字段级规则,CrossColumn:跨字段规则,Customize:自定义规则
	Type *string `json:"type,omitempty"`

	// 规则模板ID
	TemplateId *int64 `json:"template_id,omitempty"`

	// 数据连接ID
	Connection *string `json:"connection,omitempty"`

	// 数据连接类型
	ConnectionType *string `json:"connection_type,omitempty"`

	// 数据库名，当存在多个值时以逗号分隔
	Databases *string `json:"databases,omitempty"`

	// 自定义规则中的SQL脚本
	Sql *string `json:"sql,omitempty"`

	// 表名，当存在多个值时以逗号分隔
	Tables *string `json:"tables,omitempty"`

	// 对照表名，当存在多个值时以逗号分隔
	ReferenceTables *string `json:"reference_tables,omitempty"`

	// 字段名，当存在多个值时以逗号分隔
	Columns *string `json:"columns,omitempty"`

	// 对照列名，当存在多个值时以逗号分隔
	ReferenceColumns *string `json:"reference_columns,omitempty"`

	// 维度
	Dimension *string `json:"dimension,omitempty"`

	// DLI队列
	Queue *string `json:"queue,omitempty"`

	// 当规则模板为正则表达式校验时的正则表达式
	RegularExpression *string `json:"regular_expression,omitempty"`

	// 模板参数
	TemplateArguments *string `json:"template_arguments,omitempty"`

	// 规则权重
	Weight *int32 `json:"weight,omitempty"`

	// 计算范围
	CalculationRange *string `json:"calculation_range,omitempty"`

	// 计算范围SQL
	CalculationRangeSql *string `json:"calculation_range_sql,omitempty"`

	// 告警表达式
	AlarmCondition *string `json:"alarm_condition,omitempty"`
}

func (o ConsistencyRuleDetailForOpenApi) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsistencyRuleDetailForOpenApi struct{}"
	}

	return strings.Join([]string{"ConsistencyRuleDetailForOpenApi", string(data)}, " ")
}
