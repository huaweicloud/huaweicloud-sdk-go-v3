package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 数据库库级对象。
type DatabaseObject struct {

	// 该数据库在实时同步场景下的类型。取值： - config：仅当该库作为数据过滤高级设置的关联库时，需要填写，此时该库以及该库下的schemas、tables“不会”被同步到目标库，name、all属性不生效，schemas、tables需要填写被关联的相关对象。（注意：如果需要同步该库级对象，则在下级对象中填写sync_type值为config。）
	SyncType *DatabaseObjectSyncType `json:"sync_type,omitempty"`

	// 该数据库在目标库的名称（库名映射）。
	Name *string `json:"name,omitempty"`

	// 是否整库迁移或同步。注意： 1.当该库下的模式、表、列需要做数据过滤、列过滤、列映射时，填false，否则填true； 2.当该库下的表需要做附加列时，需要填true，并且在表级对象的columns里填写附加列信息； 3.当该库下的表所包含的列作为数据过滤高级设置的关联列时，需要填true，并且在columns里填写关联列信息，在config_conditions填写数据过滤高级设置的配置条件；
	All *bool `json:"all,omitempty"`

	// 需要迁移或同步的模式，当整库迁移或同步为false时需要填写。
	Schemas map[string]SchemaObject `json:"schemas,omitempty"`

	// 需要迁移或同步的表，当整库迁移或同步为false时需要填写。
	Tables map[string]TableObject `json:"tables,omitempty"`

	// 库下的表的数量，表的数量超过阈值就不显示。
	TotalTableNum *int32 `json:"total_table_num,omitempty"`

	// 是否已同步
	IsSynchronized *bool `json:"is_synchronized,omitempty"`
}

func (o DatabaseObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseObject struct{}"
	}

	return strings.Join([]string{"DatabaseObject", string(data)}, " ")
}

type DatabaseObjectSyncType struct {
	value string
}

type DatabaseObjectSyncTypeEnum struct {
	CONFIG DatabaseObjectSyncType
}

func GetDatabaseObjectSyncTypeEnum() DatabaseObjectSyncTypeEnum {
	return DatabaseObjectSyncTypeEnum{
		CONFIG: DatabaseObjectSyncType{
			value: "config",
		},
	}
}

func (c DatabaseObjectSyncType) Value() string {
	return c.value
}

func (c DatabaseObjectSyncType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatabaseObjectSyncType) UnmarshalJSON(b []byte) error {
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
