package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateEventSchemaRequest struct {

	// 指定查询的事件模型ID
	SchemaId string `json:"schema_id"`

	Body *CustomizeSchemaUpdateReq `json:"body,omitempty"`
}

func (o UpdateEventSchemaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEventSchemaRequest struct{}"
	}

	return strings.Join([]string{"UpdateEventSchemaRequest", string(data)}, " ")
}
