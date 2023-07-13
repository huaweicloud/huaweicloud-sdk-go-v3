package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentSnapshotItem struct {

	// 组件ID。
	ComponentId *string `json:"component_id,omitempty"`

	// 快照索引。
	Index *int32 `json:"index,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	Context *ComponentSnapshotContext `json:"context,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ComponentSnapshotItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentSnapshotItem struct{}"
	}

	return strings.Join([]string{"ComponentSnapshotItem", string(data)}, " ")
}
