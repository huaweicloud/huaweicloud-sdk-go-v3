package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataFilterObjectResp 数据过滤对象信息
type DataFilterObjectResp struct {

	// 数据过滤对象信息批次ID，与父级DataFilterRuleResp的batch_id关联，用于追溯对象归属的批次
	BatchId *int64 `json:"batch_id,omitempty"`

	// 数据过滤对象的唯一标识符，由对象层级结构（库名、表名、对象类型等）组合生成
	Id *string `json:"id,omitempty"`

	// 父级过滤对象的标识符，用于表示对象之间的层级归属关系（如列归属于表、表归属于库）。
	ParentId *string `json:"parent_id,omitempty"`

	// 数据过滤对象的名称，按object_type不同可为库名、表名或列名。
	ObjectName *string `json:"object_name,omitempty"`

	// 数据过滤对象的映射名称（别名），在目标端数据库对象名与源端不一致时使用。
	ObjectAliasName *string `json:"object_alias_name,omitempty"`

	// 数据过滤对象的层级类型，标识当前过滤规则粒度级别。
	ObjectType *string `json:"object_type,omitempty"`

	// 数据过滤对象所属的数据库名称。
	DbName *string `json:"db_name,omitempty"`

	// 数据过滤对象所属的schema名。
	SchemaName *string `json:"schema_name,omitempty"`

	// 数据过滤对象所属的表名。
	TableName *string `json:"table_name,omitempty"`

	// 标记该对象是否启用数据过滤。
	IsDataFilter *DataFilterObjectRespIsDataFilter `json:"is_data_filter,omitempty"`

	// 数据过滤的类型，标识使用的过滤策略。
	DataFilterType *string `json:"data_filter_type,omitempty"`

	// 数据过滤条件
	FilterCondition *string `json:"filter_condition,omitempty"`

	// 数据过滤条件是否已同步
	IsSynchronized *bool `json:"is_synchronized,omitempty"`
}

func (o DataFilterObjectResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataFilterObjectResp struct{}"
	}

	return strings.Join([]string{"DataFilterObjectResp", string(data)}, " ")
}

type DataFilterObjectRespIsDataFilter struct {
	value string
}

type DataFilterObjectRespIsDataFilterEnum struct {
	TRUE  DataFilterObjectRespIsDataFilter
	FALSE DataFilterObjectRespIsDataFilter
}

func GetDataFilterObjectRespIsDataFilterEnum() DataFilterObjectRespIsDataFilterEnum {
	return DataFilterObjectRespIsDataFilterEnum{
		TRUE: DataFilterObjectRespIsDataFilter{
			value: "true",
		},
		FALSE: DataFilterObjectRespIsDataFilter{
			value: "false",
		},
	}
}

func (c DataFilterObjectRespIsDataFilter) Value() string {
	return c.value
}

func (c DataFilterObjectRespIsDataFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataFilterObjectRespIsDataFilter) UnmarshalJSON(b []byte) error {
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
