package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEventSchemaVersionRequest struct {

	// 指定查询的事件模型ID
	SchemaId string `json:"schema_id"`

	Body *CustomizeSchemaVersionCreateReq `json:"body,omitempty"`
}

func (o CreateEventSchemaVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventSchemaVersionRequest struct{}"
	}

	return strings.Join([]string{"CreateEventSchemaVersionRequest", string(data)}, " ")
}
