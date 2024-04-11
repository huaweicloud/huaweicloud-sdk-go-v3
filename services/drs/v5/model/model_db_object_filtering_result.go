package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DbObjectFilteringResult 数据库表过滤规则响应体
type DbObjectFilteringResult struct {

	// 数据库库名。
	DbName *string `json:"db_name,omitempty"`

	// 数据库Schema名称。
	SchemaName *string `json:"schema_name,omitempty"`

	// 数据库表名称。
	TableName *string `json:"table_name,omitempty"`

	// 数据过滤校验结果。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 当数据过滤校验结果是false，返回校验失败的原因。
	Message *string `json:"message,omitempty"`

	// 对比的来源 - job 表示数据同步时的过滤 - compare 表示数据对比的过滤
	Source *DbObjectFilteringResultSource `json:"source,omitempty"`

	// 校验目标库比对条件过滤的结果
	TargetResult *string `json:"target_result,omitempty"`

	// 校验源库比对条件过滤的结果
	SourceResult *string `json:"source_result,omitempty"`

	// 校验目标库比对条件过滤的失败原因
	TargetMessage *string `json:"target_message,omitempty"`

	// 校验源库比对条件过滤的失败原因
	SourceMessage *string `json:"source_message,omitempty"`
}

func (o DbObjectFilteringResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbObjectFilteringResult struct{}"
	}

	return strings.Join([]string{"DbObjectFilteringResult", string(data)}, " ")
}

type DbObjectFilteringResultSource struct {
	value string
}

type DbObjectFilteringResultSourceEnum struct {
	JOB     DbObjectFilteringResultSource
	COMPARE DbObjectFilteringResultSource
}

func GetDbObjectFilteringResultSourceEnum() DbObjectFilteringResultSourceEnum {
	return DbObjectFilteringResultSourceEnum{
		JOB: DbObjectFilteringResultSource{
			value: "job",
		},
		COMPARE: DbObjectFilteringResultSource{
			value: "compare",
		},
	}
}

func (c DbObjectFilteringResultSource) Value() string {
	return c.value
}

func (c DbObjectFilteringResultSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DbObjectFilteringResultSource) UnmarshalJSON(b []byte) error {
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
