package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomizeSchemaVersionItemInfo struct {

	// 事件模型版本ID
	Id *string `json:"id,omitempty"`

	// 事件模型版本号
	Version *int32 `json:"version,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`
}

func (o CustomizeSchemaVersionItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeSchemaVersionItemInfo struct{}"
	}

	return strings.Join([]string{"CustomizeSchemaVersionItemInfo", string(data)}, " ")
}
