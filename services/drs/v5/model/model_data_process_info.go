package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataProcessInfo struct {

	// 指定任务数据加工规则请求体
	FilterConditions *[]DataFilteringCondition `json:"filter_conditions,omitempty"`

	// 库级、批量表级处理为true，单表操作为false
	IsBatchProcess *bool `json:"is_batch_process,omitempty"`

	// 附加列 当选择附加列时必须填写 约束：使用多对一操作时，需要使用数据加工的附加列操作来避免数据冲突。
	AddColumns *[]AddColumnInfo `json:"add_columns,omitempty"`

	// 支持DDL的语法 选择增量迁移或同步的DDL操作。取值及意思如下： \"table\": \"CREATE TABLE, ALTER TABLE,DROP TABLE,RENAME TABLE\" 如该值为空，不迁移或同步DDL操作
	DdlOperation map[string]string `json:"ddl_operation,omitempty"`

	// 支持DML的语法 选择DML操作时，取值如下： i：INSERT。 u：UPDATE。 d：DELETE。 如该值为空，不增量迁移或同步DML操作。
	DmlOperation *string `json:"dml_operation,omitempty"`

	DbObjectColumnInfo *DbObjectColumnInfo `json:"db_object_column_info,omitempty"`

	DbOrTableRenameRule *DbOrTableRenameRule `json:"db_or_table_rename_rule,omitempty"`

	DbObject *DbObject `json:"db_object,omitempty"`

	// 表示该规则是否已同步至目标库
	IsSynchronized *bool `json:"is_synchronized,omitempty"`

	// 对比的来源 - job 表示数据同步时的过滤 - compare 表示数据对比的过滤
	Source *DataProcessInfoSource `json:"source,omitempty"`
}

func (o DataProcessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataProcessInfo struct{}"
	}

	return strings.Join([]string{"DataProcessInfo", string(data)}, " ")
}

type DataProcessInfoSource struct {
	value string
}

type DataProcessInfoSourceEnum struct {
	JOB     DataProcessInfoSource
	COMPARE DataProcessInfoSource
}

func GetDataProcessInfoSourceEnum() DataProcessInfoSourceEnum {
	return DataProcessInfoSourceEnum{
		JOB: DataProcessInfoSource{
			value: "job",
		},
		COMPARE: DataProcessInfoSource{
			value: "compare",
		},
	}
}

func (c DataProcessInfoSource) Value() string {
	return c.value
}

func (c DataProcessInfoSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataProcessInfoSource) UnmarshalJSON(b []byte) error {
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
