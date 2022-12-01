package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 数据库表级对象。
type TableObject struct {

	// 该表在实时同步场景下的类型。取值： - config：仅当该表作为数据过滤高级设置的关联表时，需要填写，此时该表以及该表下的columns“不会”被同步到目标库，name、all、filtered、filter_conditions属性不生效，columns需要填写被关联的相关对象，config_conditions需要填写数据过滤高级设置的配置条件。（注意：如果需要同步该表级对象，则在下级对象中填写sync_type值为config。）
	SyncType *string `json:"sync_type,omitempty"`

	// 对象类型。取值： - table：表。 - view：视图。 - procedure：存储过程。
	Type *TableObjectType `json:"type,omitempty"`

	// 该表在目标库的名称（表名映射）。
	Name *string `json:"name,omitempty"`

	// 是否整表迁移或同步。注意： 1.当该表不需要做列过滤、列映射时，填true，如果需要做列过滤、列映射则填false； 2.当该表需要做附加列时，需要填true，并且在columns里填写附加列信息； 3.当该表所包含的列作为数据过滤高级设置的关联列时，需要填true，并且在columns里填写关联列信息、config_conditions填写数据过滤高级设置的配置条件；
	All *bool `json:"all,omitempty"`

	// 一对多情况下，表级上对库名的映射。
	DbAliasName *string `json:"db_alias_name,omitempty"`

	// 一对多情况下，表级上对schema名的映射。
	SchemaAliasName *string `json:"schema_alias_name,omitempty"`

	// 该表是否进行数据过滤。
	Filtered *bool `json:"filtered,omitempty"`

	// 该表数据的过滤条件，生成加工规则值为SQL条件语句，长度限制512。
	FilterConditions *[]string `json:"filter_conditions,omitempty"`

	// 该表数据过滤高级设置的配置条件，当该表作为联表查询时填写，生成加工规则值为SQL条件语句，长度限制512。
	ConfigConditions *[]string `json:"config_conditions,omitempty"`

	// 是否已经进行同步。
	HasBeenSent *bool `json:"has_been_sent,omitempty"`

	// 需要同步/映射/过滤/新增的列，当需要列过滤、列映射、附加列功能时填写，仅在实时同步任务中生效，当整表同步为false时需要填写。
	Columns map[string]ColumnObject `json:"columns,omitempty"`
}

func (o TableObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableObject struct{}"
	}

	return strings.Join([]string{"TableObject", string(data)}, " ")
}

type TableObjectType struct {
	value string
}

type TableObjectTypeEnum struct {
	TABLE     TableObjectType
	VIEW      TableObjectType
	PROCEDURE TableObjectType
}

func GetTableObjectTypeEnum() TableObjectTypeEnum {
	return TableObjectTypeEnum{
		TABLE: TableObjectType{
			value: "table",
		},
		VIEW: TableObjectType{
			value: "view",
		},
		PROCEDURE: TableObjectType{
			value: "procedure",
		},
	}
}

func (c TableObjectType) Value() string {
	return c.value
}

func (c TableObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableObjectType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
