package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QualityTaskRuleDetailForOpenApi struct {

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

	// 自定义规则中的SQL脚本，系统内置规则时返回
	Sql *string `json:"sql,omitempty"`

	// 表名，当存在多个值时以逗号分隔
	Tables *string `json:"tables,omitempty"`

	// 跨字段规则中的参考表名，当存在多个值时以逗号分隔
	ReferenceTables *string `json:"reference_tables,omitempty"`

	// 字段名，当存在多个值时以逗号分隔
	Columns *string `json:"columns,omitempty"`

	// 跨字段规则中的参考字段名，当存在多个值时以逗号分隔
	ReferenceColumns *string `json:"reference_columns,omitempty"`

	// 是否忽视规则错误
	IgnoreError *bool `json:"ignore_error,omitempty"`

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

	// 是否导出异常数据
	ExportAbnormalTable *bool `json:"export_abnormal_table,omitempty"`

	// 异常表数据库
	AbnormalTableDatabase *string `json:"abnormal_table_database,omitempty"`

	// 异常表Schema
	AbnormalTableSchema *string `json:"abnormal_table_schema,omitempty"`

	// 异常字段所在的表
	AbnormalTable *string `json:"abnormal_table,omitempty"`

	// 异常表前缀
	AbnormalTablePrefix *string `json:"abnormal_table_prefix,omitempty"`

	// 异常表后缀
	AbnormalTableSuffix *string `json:"abnormal_table_suffix,omitempty"`

	// 异常字段名，当存在多个值时以逗号分隔
	AbnormalTableColumns *string `json:"abnormal_table_columns,omitempty"`

	// 异常表SQL
	AbnormalTableSql *string `json:"abnormal_table_sql,omitempty"`

	// 异常表是否输出规则配置
	AbnormalTableOutConfig *bool `json:"abnormal_table_out_config,omitempty"`

	// 异常表是否包含空值
	AbnormalTableIncludeNullValue *bool `json:"abnormal_table_include_null_value,omitempty"`

	// 异常表输出行数, 0代表全量输出
	AbnormalTableOutDataNumber *int32 `json:"abnormal_table_out_data_number,omitempty"`

	// 是否开启质量评分
	ScoreSwitch *bool `json:"score_switch,omitempty"`

	// 质量评分表所在schema
	ScoreSchema *string `json:"score_schema,omitempty"`

	// 质量评分表名
	ScoreTable *string `json:"score_table,omitempty"`

	// 质量评分表达式
	ScoreExpression *string `json:"score_expression,omitempty"`
}

func (o QualityTaskRuleDetailForOpenApi) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QualityTaskRuleDetailForOpenApi struct{}"
	}

	return strings.Join([]string{"QualityTaskRuleDetailForOpenApi", string(data)}, " ")
}
