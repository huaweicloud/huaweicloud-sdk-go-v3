package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemInfo 关联需求
type WorkItemInfo struct {

	// 工作项编号
	WorkItemId string `json:"work_item_id"`

	// 是否有子需求
	HasChild bool `json:"has_child"`

	// 是否展开
	IsOpen bool `json:"is_open"`

	// 子需求
	ChildList *[]WorkItemInfo `json:"child_list,omitempty"`
}

func (o WorkItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemInfo struct{}"
	}

	return strings.Join([]string{"WorkItemInfo", string(data)}, " ")
}
