package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventSchemaVersionResponse Response Object
type CreateEventSchemaVersionResponse struct {

	// 事件模型版本ID
	Id *string `json:"id,omitempty"`

	// 事件模型ID
	SchemaId *string `json:"schema_id,omitempty"`

	// 事件模型版本号
	Version *int32 `json:"version,omitempty"`

	// 事件模型格式
	Format *string `json:"format,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 事件模型内容定义
	Definition *string `json:"definition,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEventSchemaVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventSchemaVersionResponse struct{}"
	}

	return strings.Join([]string{"CreateEventSchemaVersionResponse", string(data)}, " ")
}
