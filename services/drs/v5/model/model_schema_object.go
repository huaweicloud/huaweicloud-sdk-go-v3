package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库模式对象。
type SchemaObject struct {

	// 该模式在实时同步场景下的类型。取值： - config：仅当该模式作为数据过滤高级设置的关联模式时，需要填写，此时该模式以及该模式下的tables“不会”被同步到目标库，name、all属性不生效，tables需要填写被关联的相关对象。（注意：如果需要同步该模式对象，则在下级对象中填写sync_type值为config。）
	SyncType *string `json:"sync_type,omitempty"`

	// 该模式在目标库的名称（模式名映射）。
	Name *string `json:"name,omitempty"`

	// 是否整模式迁移或同步。注意： 1.当该模式下的表、列需要做数据过滤、列过滤、列映射时，填false，否则填true； 2.当该模式下的表需要做附加列时，需要填true，并且在表级对象的columns里填写附加列信息； 3.当该模式下的表所包含的列作为数据过滤高级设置的关联列时，需要填true，并且在columns里填写关联列信息、config_conditions填写数据过滤高级设置的配置条件；
	All *bool `json:"all,omitempty"`

	// 需要迁移或同步的表，当整模式迁移或同步为false时需要填写。
	Tables map[string]TableObject `json:"tables,omitempty"`
}

func (o SchemaObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchemaObject struct{}"
	}

	return strings.Join([]string{"SchemaObject", string(data)}, " ")
}
