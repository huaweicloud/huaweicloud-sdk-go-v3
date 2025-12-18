package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListTableVolumesRequestBody struct {

	// **参数解释**: 数据库名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	DatabaseName string `json:"database_name"`

	// **参数解释**: schema名称。 **约束限制**: 不涉及。
	SchemaNames []string `json:"schema_names"`

	// **参数解释**: 表名称。。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	TableName *string `json:"table_name,omitempty"`

	// **参数解释**: 表所属用户名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 排序字段。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SortKey *ListTableVolumesRequestBodySortKey `json:"sort_key,omitempty"`

	// **参数解释**: 排序方法。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SortOrder *ListTableVolumesRequestBodySortOrder `json:"sort_order,omitempty"`

	// **参数解释**: 查询记录数 **约束限制**: 必须为数字，不能为负数。 **取值范围**: 1 - 100 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。 **约束限制**: 必须为数字，不能为负数。 **取值范围**: 0 - 10000 **默认取值**: 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListTableVolumesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableVolumesRequestBody struct{}"
	}

	return strings.Join([]string{"ListTableVolumesRequestBody", string(data)}, " ")
}

type ListTableVolumesRequestBodySortKey struct {
	value string
}

type ListTableVolumesRequestBodySortKeyEnum struct {
	TABLE_SIZE          ListTableVolumesRequestBodySortKey
	ID                  ListTableVolumesRequestBodySortKey
	TABLE_NAME          ListTableVolumesRequestBodySortKey
	TABLE_OWNER         ListTableVolumesRequestBodySortKey
	DATABASE_NAME       ListTableVolumesRequestBodySortKey
	SCHEMA_NAME         ListTableVolumesRequestBodySortKey
	IS_PART_TYPE        ListTableVolumesRequestBodySortKey
	IS_HASH_CLUSTER_KEY ListTableVolumesRequestBodySortKey
	TUPLES              ListTableVolumesRequestBodySortKey
	CREATE_TIME         ListTableVolumesRequestBodySortKey
	UPDATE_TIME         ListTableVolumesRequestBodySortKey
	AVERAGE_SIZE        ListTableVolumesRequestBodySortKey
	MAX_RATIO           ListTableVolumesRequestBodySortKey
	MIN_RATIO           ListTableVolumesRequestBodySortKey
	SKEW_SIZE           ListTableVolumesRequestBodySortKey
	SKEW_RATIO          ListTableVolumesRequestBodySortKey
	SKEW_STDDEV         ListTableVolumesRequestBodySortKey
}

func GetListTableVolumesRequestBodySortKeyEnum() ListTableVolumesRequestBodySortKeyEnum {
	return ListTableVolumesRequestBodySortKeyEnum{
		TABLE_SIZE: ListTableVolumesRequestBodySortKey{
			value: "table_size",
		},
		ID: ListTableVolumesRequestBodySortKey{
			value: "id",
		},
		TABLE_NAME: ListTableVolumesRequestBodySortKey{
			value: "table_name",
		},
		TABLE_OWNER: ListTableVolumesRequestBodySortKey{
			value: "table_owner",
		},
		DATABASE_NAME: ListTableVolumesRequestBodySortKey{
			value: "database_name",
		},
		SCHEMA_NAME: ListTableVolumesRequestBodySortKey{
			value: "schema_name",
		},
		IS_PART_TYPE: ListTableVolumesRequestBodySortKey{
			value: "is_part_type",
		},
		IS_HASH_CLUSTER_KEY: ListTableVolumesRequestBodySortKey{
			value: "is_hash_cluster_key",
		},
		TUPLES: ListTableVolumesRequestBodySortKey{
			value: "tuples",
		},
		CREATE_TIME: ListTableVolumesRequestBodySortKey{
			value: "create_time",
		},
		UPDATE_TIME: ListTableVolumesRequestBodySortKey{
			value: "update_time",
		},
		AVERAGE_SIZE: ListTableVolumesRequestBodySortKey{
			value: "average_size",
		},
		MAX_RATIO: ListTableVolumesRequestBodySortKey{
			value: "max_ratio",
		},
		MIN_RATIO: ListTableVolumesRequestBodySortKey{
			value: "min_ratio",
		},
		SKEW_SIZE: ListTableVolumesRequestBodySortKey{
			value: "skew_size",
		},
		SKEW_RATIO: ListTableVolumesRequestBodySortKey{
			value: "skew_ratio",
		},
		SKEW_STDDEV: ListTableVolumesRequestBodySortKey{
			value: "skew_stddev",
		},
	}
}

func (c ListTableVolumesRequestBodySortKey) Value() string {
	return c.value
}

func (c ListTableVolumesRequestBodySortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTableVolumesRequestBodySortKey) UnmarshalJSON(b []byte) error {
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

type ListTableVolumesRequestBodySortOrder struct {
	value string
}

type ListTableVolumesRequestBodySortOrderEnum struct {
	DESC ListTableVolumesRequestBodySortOrder
	ASC  ListTableVolumesRequestBodySortOrder
}

func GetListTableVolumesRequestBodySortOrderEnum() ListTableVolumesRequestBodySortOrderEnum {
	return ListTableVolumesRequestBodySortOrderEnum{
		DESC: ListTableVolumesRequestBodySortOrder{
			value: "DESC",
		},
		ASC: ListTableVolumesRequestBodySortOrder{
			value: "ASC",
		},
	}
}

func (c ListTableVolumesRequestBodySortOrder) Value() string {
	return c.value
}

func (c ListTableVolumesRequestBodySortOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTableVolumesRequestBodySortOrder) UnmarshalJSON(b []byte) error {
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
