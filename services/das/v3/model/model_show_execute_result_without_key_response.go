package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExecuteResultWithoutKeyResponse Response Object
type ShowExecuteResultWithoutKeyResponse struct {

	// 字段值
	ColumnValues *[]interface{} `json:"column_values,omitempty"`

	// 字段属性，字段名等
	Metadata *[]Column `json:"metadata,omitempty"`

	// 结果集类型
	ResultType *string `json:"result_type,omitempty"`

	// 行数
	Rows *int32 `json:"rows,omitempty"`

	// 执行耗时
	TimeDelay *int64 `json:"time_delay,omitempty"`

	// 结果集是否可编辑
	CanEdit *bool `json:"can_edit,omitempty"`

	// 结果集是否可导出
	CanExport *bool `json:"can_export,omitempty"`

	// 编辑库名
	EditDbName *string `json:"edit_db_name,omitempty"`

	// 编辑表名
	EditTable *string `json:"edit_table,omitempty"`

	// 主键信息
	EditPrimaryKeys *[]string `json:"edit_primary_keys,omitempty"`

	// 不能编辑的原因
	CannotEditReason *int32 `json:"cannot_edit_reason,omitempty"`

	// 额外信息
	ExtendDatas *[]interface{} `json:"extend_datas,omitempty"`

	// 数据总量
	DataSum *int32 `json:"data_sum,omitempty"`

	// 是否为大表
	BigTable *bool `json:"big_table,omitempty"`

	// 执行命令告警信息
	Warning *[]string `json:"warning,omitempty"`

	// 要执行的SQL语句
	Sql *string `json:"sql,omitempty"`

	// 是否为执行计划语句
	ExplainSql *bool `json:"explain_sql,omitempty"`

	// 页面状态
	PageState *string `json:"page_state,omitempty"`

	// 查询结果是否超过规定大小
	ExceedData *bool `json:"exceed_data,omitempty"`

	// 执行状态（finished：执行完毕，pending：执行中）
	ExecuteStatus  *string `json:"execute_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowExecuteResultWithoutKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExecuteResultWithoutKeyResponse struct{}"
	}

	return strings.Join([]string{"ShowExecuteResultWithoutKeyResponse", string(data)}, " ")
}
