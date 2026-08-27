package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelConfigFailedItem 关联操作失败项。
type ModelConfigFailedItem struct {

	// 模型分组ID。
	GroupId *string `json:"group_id,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 失败原因。
	Reason *string `json:"reason,omitempty"`
}

func (o ModelConfigFailedItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelConfigFailedItem struct{}"
	}

	return strings.Join([]string{"ModelConfigFailedItem", string(data)}, " ")
}
