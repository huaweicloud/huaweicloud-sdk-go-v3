package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailOfEventSchemaRequest struct {

	// 指定查询的事件模型ID
	SchemaId string `json:"schema_id"`
}

func (o ShowDetailOfEventSchemaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfEventSchemaRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailOfEventSchemaRequest", string(data)}, " ")
}
