package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEventSchemaVersionRequest struct {

	// 指定查询的事件模型ID
	SchemaId string `json:"schema_id"`

	// 指定查询的事件模型版本号
	Version int32 `json:"version"`
}

func (o DeleteEventSchemaVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventSchemaVersionRequest struct{}"
	}

	return strings.Join([]string{"DeleteEventSchemaVersionRequest", string(data)}, " ")
}
