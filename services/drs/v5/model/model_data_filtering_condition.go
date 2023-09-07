package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataFilteringCondition 数据过滤数条件信息体
type DataFilteringCondition struct {

	// 过滤条件 当filtering_type是configConditionalFilter时， value默认填写config 当filtering_type是contentConditionalFilter时， value默认填写过滤条件  注意： 每张表仅支持添加一个校验规则。 数据过滤每次最多支持500张表。 过滤表达式不支持使用某种数据库引擎特有的package、函数、变量、常量等写法，须使用通用SQL标准。请直接输入SQL语句中WHERE之后的部分（不包含WHERE和分号，例如：sid > 3 and sname like \"G %\"），最多支持输入512个字符。 过滤条件填写的SQL语句中，关键字需要用反引号，datatime类型（包含日期和时间）需要用单引号，例如：update > '2022-07-13 00:00:00' and age >10。 不支持对LOB字段设置过滤条件，如CLOB、BLOB、BYTEA等大字段类型。 不支持库名、表名带有换行符的对象设置过滤规则。 建议不要对非精确类型字段设置过滤条件，如FLOAT、DECIMAL、DOUBLE等。 建议不要对带有特殊字符的字段设置过滤条件。 不建议使用非幂等表达式或函数作为数据加工条件，如SYSTIMESTAMP，SYSDATE等，因其每次调用返回的结果可能会有差异，导致达不到预期。
	Value *string `json:"value,omitempty"`

	// 过滤条件类型  contentConditionalFilter: 简单条件过滤 configConditionalFilter: 关联表过滤
	FilteringType *DataFilteringConditionFilteringType `json:"filtering_type,omitempty"`
}

func (o DataFilteringCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataFilteringCondition struct{}"
	}

	return strings.Join([]string{"DataFilteringCondition", string(data)}, " ")
}

type DataFilteringConditionFilteringType struct {
	value string
}

type DataFilteringConditionFilteringTypeEnum struct {
	CONTENT_CONDITIONAL_FILTER DataFilteringConditionFilteringType
	CONFIG_CONDITIONAL_FILTER  DataFilteringConditionFilteringType
}

func GetDataFilteringConditionFilteringTypeEnum() DataFilteringConditionFilteringTypeEnum {
	return DataFilteringConditionFilteringTypeEnum{
		CONTENT_CONDITIONAL_FILTER: DataFilteringConditionFilteringType{
			value: "contentConditionalFilter",
		},
		CONFIG_CONDITIONAL_FILTER: DataFilteringConditionFilteringType{
			value: " configConditionalFilter",
		},
	}
}

func (c DataFilteringConditionFilteringType) Value() string {
	return c.value
}

func (c DataFilteringConditionFilteringType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataFilteringConditionFilteringType) UnmarshalJSON(b []byte) error {
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
