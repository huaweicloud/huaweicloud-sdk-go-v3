package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExportTableVolumeRequestBody struct {

	// **参数解释**: 数据库名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	DatabaseName string `json:"database_name"`

	// **参数解释**: schema名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SchemaNames []string `json:"schema_names"`

	// **参数解释**: 表名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	TableName *string `json:"table_name,omitempty"`

	// **参数解释**: 表所属用户名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 排序字段。 **约束限制**: 不涉及。 **取值范围**: - table_size：表的大小。 - id：表ID。 - table_name：表名称。 - table_owner：表所属用户名称。 - database_name：数据库名称。 - schema_name：schema名称。 - is_part_type：表或者索引是否具有分区表的性质。 - is_hash_cluster_key：是否包含hash分区列信息。 - tuples：表中行的数目。 - create_time：创建时间。 - update_time：修改时间。 - average_size：表大小平均值(totalsize/DN个数，该值为平均分布的理想情况下，表在各DN占用空间大小)。 - max_ratio：单DN表大小最大值占比（表在各DN占用空间的最大值/totalsize）。 - min_ratio：单DN表大小最小值占比（表在各DN占用空间的最小值/totalsize）。 - skew_size：表分布倾斜值（单DN表大小最大值 - 单DN表大小最小值）。 - skew_ratio：表分布倾斜率（skewsize/totalsize）。 - skew_stddev：表分布标准方差（在表大小一定的情况下，该值越大表明表的整体分布情况越倾斜）。  **默认取值** 不涉及。
	SortKey *ExportTableVolumeRequestBodySortKey `json:"sort_key,omitempty"`

	// **参数解释** 实时会话统计排序方式。 **约束限制**: 不涉及。 **取值范围** - ASC：根据sort_key值升序。 - DESC：根据sort_key值降序。  **默认取值** ASC
	SortOrder *ExportTableVolumeRequestBodySortOrder `json:"sort_order,omitempty"`
}

func (o ExportTableVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTableVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"ExportTableVolumeRequestBody", string(data)}, " ")
}

type ExportTableVolumeRequestBodySortKey struct {
	value string
}

type ExportTableVolumeRequestBodySortKeyEnum struct {
	TABLE_SIZE          ExportTableVolumeRequestBodySortKey
	ID                  ExportTableVolumeRequestBodySortKey
	TABLE_NAME          ExportTableVolumeRequestBodySortKey
	TABLE_OWNER         ExportTableVolumeRequestBodySortKey
	DATABASE_NAME       ExportTableVolumeRequestBodySortKey
	SCHEMA_NAME         ExportTableVolumeRequestBodySortKey
	IS_PART_TYPE        ExportTableVolumeRequestBodySortKey
	IS_HASH_CLUSTER_KEY ExportTableVolumeRequestBodySortKey
	TUPLES              ExportTableVolumeRequestBodySortKey
	CREATE_TIME         ExportTableVolumeRequestBodySortKey
	UPDATE_TIME         ExportTableVolumeRequestBodySortKey
	AVERAGE_SIZE        ExportTableVolumeRequestBodySortKey
	MAX_RATIO           ExportTableVolumeRequestBodySortKey
	MIN_RATIO           ExportTableVolumeRequestBodySortKey
	SKEW_SIZE           ExportTableVolumeRequestBodySortKey
	SKEW_RATIO          ExportTableVolumeRequestBodySortKey
	SKEW_STDDEV         ExportTableVolumeRequestBodySortKey
}

func GetExportTableVolumeRequestBodySortKeyEnum() ExportTableVolumeRequestBodySortKeyEnum {
	return ExportTableVolumeRequestBodySortKeyEnum{
		TABLE_SIZE: ExportTableVolumeRequestBodySortKey{
			value: "table_size",
		},
		ID: ExportTableVolumeRequestBodySortKey{
			value: "id",
		},
		TABLE_NAME: ExportTableVolumeRequestBodySortKey{
			value: "table_name",
		},
		TABLE_OWNER: ExportTableVolumeRequestBodySortKey{
			value: "table_owner",
		},
		DATABASE_NAME: ExportTableVolumeRequestBodySortKey{
			value: "database_name",
		},
		SCHEMA_NAME: ExportTableVolumeRequestBodySortKey{
			value: "schema_name",
		},
		IS_PART_TYPE: ExportTableVolumeRequestBodySortKey{
			value: "is_part_type",
		},
		IS_HASH_CLUSTER_KEY: ExportTableVolumeRequestBodySortKey{
			value: "is_hash_cluster_key",
		},
		TUPLES: ExportTableVolumeRequestBodySortKey{
			value: "tuples",
		},
		CREATE_TIME: ExportTableVolumeRequestBodySortKey{
			value: "create_time",
		},
		UPDATE_TIME: ExportTableVolumeRequestBodySortKey{
			value: "update_time",
		},
		AVERAGE_SIZE: ExportTableVolumeRequestBodySortKey{
			value: "average_size",
		},
		MAX_RATIO: ExportTableVolumeRequestBodySortKey{
			value: "max_ratio",
		},
		MIN_RATIO: ExportTableVolumeRequestBodySortKey{
			value: "min_ratio",
		},
		SKEW_SIZE: ExportTableVolumeRequestBodySortKey{
			value: "skew_size",
		},
		SKEW_RATIO: ExportTableVolumeRequestBodySortKey{
			value: "skew_ratio",
		},
		SKEW_STDDEV: ExportTableVolumeRequestBodySortKey{
			value: "skew_stddev",
		},
	}
}

func (c ExportTableVolumeRequestBodySortKey) Value() string {
	return c.value
}

func (c ExportTableVolumeRequestBodySortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportTableVolumeRequestBodySortKey) UnmarshalJSON(b []byte) error {
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

type ExportTableVolumeRequestBodySortOrder struct {
	value string
}

type ExportTableVolumeRequestBodySortOrderEnum struct {
	ASC  ExportTableVolumeRequestBodySortOrder
	DESC ExportTableVolumeRequestBodySortOrder
}

func GetExportTableVolumeRequestBodySortOrderEnum() ExportTableVolumeRequestBodySortOrderEnum {
	return ExportTableVolumeRequestBodySortOrderEnum{
		ASC: ExportTableVolumeRequestBodySortOrder{
			value: "ASC",
		},
		DESC: ExportTableVolumeRequestBodySortOrder{
			value: "DESC",
		},
	}
}

func (c ExportTableVolumeRequestBodySortOrder) Value() string {
	return c.value
}

func (c ExportTableVolumeRequestBodySortOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportTableVolumeRequestBodySortOrder) UnmarshalJSON(b []byte) error {
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
