package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库列信息。
type ColumnObject struct {

	// 该列在实时同步场景下的类型。取值： - config：当该列作为数据过滤高级设置的关联列时，需要填写，同时如果该列是主建或优化查询所需的索引，则需要填写primary_key_for_data_filtering或index_for_data_filtering。（注意：是否同步该列到目标库由“filtered”属性控制，与库级、模式级、表级控制方式不同。）
	SyncType *string `json:"sync_type,omitempty"`

	// 该列是否在数据过滤高级设置场景下为主键，如果是主建则填该列列名，否则不填。
	PrimaryKeyForDataFiltering *string `json:"primary_key_for_data_filtering,omitempty"`

	// 优化查询所需的索引，将会为缓存数据增加索引，不会影响源表，当该列作为数据过滤高级设置的关联索引时，需要填写，否则不填。
	IndexForDataFiltering *string `json:"index_for_data_filtering,omitempty"`

	// 该列在目标库的名称（列名映射），当该列为“附加列”时须与数据库表级对象中列名保持一致。
	Name *string `json:"name,omitempty"`

	// 该列字段的数据类型。 列过滤：填写源列字段的数据类型。 附加列：新填充的列指定字段的数据类型，根据不同操作类型来决定取值范围与约束。取值： - 以默认值方式，支持：int,long,varchar(256),datetime,timestamp。 - 以create_time为列，支持：long,datetime,timestamp。 - 以update_time为列，支持：long,datetime,timestamp。 - 以表达式为列，支持：varchar(256)，且列值仅为：concat(__current_database, '@', __current_table) - 以serverName@database@table为列，支持：varchar(256)。
	Type *string `json:"type,omitempty"`

	// 该列是否在列映射场景下为主键，如果是主建则填PRI，否则填空。
	PrimaryKeyForColumnFiltering *string `json:"primary_key_for_column_filtering,omitempty"`

	// 该列是否进列过滤，不能与附加列additional同时使用。取值： - true：表示同步该列。 - false：表示过滤该列不同步。
	Filtered *bool `json:"filtered,omitempty"`

	// 该列是否为附加列，当该列为附加列时：name必须与表级对象中列名一致，并且不能与列过滤filtered同时使用。
	Additional *bool `json:"additional,omitempty"`

	// 操作类型，以特定的操作类型填充新加的列。取值： - 以默认值方式：\"operation_type\":\"ADDITIONALCOLUMN,default_value\" - 以create_time为列：\"operation_type\":\"ADDITIONALCOLUMN,create_time\" - 以update_time为列：\"operation_type\":\"ADDITIONALCOLUMN,update_time\" - 以表达式为列：\"operation_type\":\"ADDITIONALCOLUMN,expression\" - 以serverName@database@table为列：\"operation_type\":\"ADDITIONALCOLUMN,server_database_table\"
	OperationType *string `json:"operation_type,omitempty"`

	// 附加列的值。约束： - 当操作类型仅“以默认值方式”，“以serverName@database@table为列”时，才支持输入对应字段类型的值。 - 当操作类型为“以表达式为列”时，该字段为固定值\"concat(__current_database, '@', __current_table)\"，不需要填写。
	Value *string `json:"value,omitempty"`
}

func (o ColumnObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnObject struct{}"
	}

	return strings.Join([]string{"ColumnObject", string(data)}, " ")
}
