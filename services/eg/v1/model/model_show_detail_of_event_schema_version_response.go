package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDetailOfEventSchemaVersionResponse struct {

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

	// 事件示例
	DataSample *string `json:"data_sample,omitempty"`

	// 事件模型内容定义
	Definition     *string `json:"definition,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDetailOfEventSchemaVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfEventSchemaVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowDetailOfEventSchemaVersionResponse", string(data)}, " ")
}
