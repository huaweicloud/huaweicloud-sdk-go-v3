package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtraFieldInfo struct {

	// 扩展字段
	CurrentCloudServiceId *string `json:"current_cloud_service_id,omitempty"`

	// 扩展字段
	Description *string `json:"description,omitempty"`

	// 扩展字段
	LevelId *string `json:"level_id,omitempty"`

	// 扩展字段
	MtmRegion *string `json:"mtm_region,omitempty"`

	// 扩展字段
	MtmType *string `json:"mtm_type,omitempty"`

	// 扩展字段
	SourceId *string `json:"source_id,omitempty"`

	// 扩展字段
	Title *string `json:"title,omitempty"`

	// 是否变更事件
	IsChangeEvent *bool `json:"is_change_event,omitempty"`

	// 是否变更事件
	IsServiceInterrupt *bool `json:"is_service_interrupt,omitempty"`
}

func (o ExtraFieldInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtraFieldInfo struct{}"
	}

	return strings.Join([]string{"ExtraFieldInfo", string(data)}, " ")
}
