package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentSnapshotItem struct {

	// 组件id。
	ComponentId *string `json:"component_id,omitempty"`

	// 上下文信息。
	Context *interface{} `json:"context,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 快照索引。
	Index *int32 `json:"index,omitempty"`

	// 修改时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ComponentSnapshotItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentSnapshotItem struct{}"
	}

	return strings.Join([]string{"ComponentSnapshotItem", string(data)}, " ")
}
