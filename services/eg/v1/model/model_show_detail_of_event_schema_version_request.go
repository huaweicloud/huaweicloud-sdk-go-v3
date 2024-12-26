package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailOfEventSchemaVersionRequest Request Object
type ShowDetailOfEventSchemaVersionRequest struct {

	// 指定查询的事件模型ID
	SchemaId string `json:"schema_id"`

	// 指定查询的事件模型版本号
	Version int32 `json:"version"`
}

func (o ShowDetailOfEventSchemaVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfEventSchemaVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailOfEventSchemaVersionRequest", string(data)}, " ")
}
