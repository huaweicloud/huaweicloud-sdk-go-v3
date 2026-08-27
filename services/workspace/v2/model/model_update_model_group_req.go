package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelGroupReq 更新模型组请求。
type UpdateModelGroupReq struct {

	// 分组名称。
	Name *string `json:"name,omitempty"`

	// 分组描述。
	Description *string `json:"description,omitempty"`

	// 分组优先级，最小值为1。
	Priority *int32 `json:"priority,omitempty"`

	// 默认模型ID。
	DefaultModelId *string `json:"default_model_id,omitempty"`
}

func (o UpdateModelGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelGroupReq struct{}"
	}

	return strings.Join([]string{"UpdateModelGroupReq", string(data)}, " ")
}
