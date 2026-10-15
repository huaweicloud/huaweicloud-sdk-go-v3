package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataFilterObjectInfo 数据过滤对象信息
type DataFilterObjectInfo struct {

	// 数据过滤对象的唯一标识。用于标识待过滤的数据库对象，填写格式为层级路径形式。 约束：填写格式为以下三种之一：库名-*-*-表名、schema名-*-*-表名、库名-*-*-schema名-*-*-表名。
	Id string `json:"id"`

	// 数据过滤对象的父级对象标识。用于标识当前过滤对象所属的上级数据库对象（库或Schema）。 约束：填写格式为以下三种之一：库名、schema名、库名-*-*-schema名。
	ParentId string `json:"parent_id"`

	// 数据过滤对象的名称。即当前需要进行过滤操作的具体数据库表名。
	ObjectName string `json:"object_name"`

	// 数据过滤对象映射名称
	ObjectAliasName *string `json:"object_alias_name,omitempty"`

	// 数据过滤对象的类型。标识当前过滤对象的数据库对象类型，如表、视图、存储过程等。
	ObjectType *string `json:"object_type,omitempty"`
}

func (o DataFilterObjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataFilterObjectInfo struct{}"
	}

	return strings.Join([]string{"DataFilterObjectInfo", string(data)}, " ")
}
