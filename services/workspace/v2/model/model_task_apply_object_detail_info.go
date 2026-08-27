package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskApplyObjectDetailInfo 任务应用对象名称信息
type TaskApplyObjectDetailInfo struct {

	// 对象id
	ObjectId *string `json:"object_id,omitempty"`

	// 对象类型
	ObjectType *string `json:"object_type,omitempty"`

	// 对象名称
	ObjectName *string `json:"object_name,omitempty"`

	// 对象desktopId
	ObjectExtraId *string `json:"object_extra_id,omitempty"`
}

func (o TaskApplyObjectDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskApplyObjectDetailInfo struct{}"
	}

	return strings.Join([]string{"TaskApplyObjectDetailInfo", string(data)}, " ")
}
