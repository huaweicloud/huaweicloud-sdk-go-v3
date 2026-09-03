package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDbObjNewRequest Request Object
type DeleteDbObjNewRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 数据库名称
	DbName string `json:"db_name"`

	// Schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名
	TableName *string `json:"table_name,omitempty"`

	// 对象名称
	ObjName *string `json:"obj_name,omitempty"`

	// 对象ID
	ObjId *string `json:"obj_id,omitempty"`

	// 对象子类型
	ObjectSubType *string `json:"object_sub_type,omitempty"`

	// 对象类型
	ObjType string `json:"obj_type"`
}

func (o DeleteDbObjNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbObjNewRequest struct{}"
	}

	return strings.Join([]string{"DeleteDbObjNewRequest", string(data)}, " ")
}
