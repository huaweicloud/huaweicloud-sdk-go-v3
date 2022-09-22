package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEventSchemaRequest struct {

	// 指定查询的事件模型ID
	SchemaId string `json:"schema_id"`
}

func (o DeleteEventSchemaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventSchemaRequest struct{}"
	}

	return strings.Join([]string{"DeleteEventSchemaRequest", string(data)}, " ")
}
