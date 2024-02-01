package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecutionPlanSummary struct {

	// 新增资源数
	ResourceAdd *int32 `json:"resource_add,omitempty"`

	// 更新资源数
	ResourceUpdate *int32 `json:"resource_update,omitempty"`

	// 删除资源数
	ResourceDelete *int32 `json:"resource_delete,omitempty"`

	// 导入资源数
	ResourceImport *int32 `json:"resource_import,omitempty"`
}

func (o ExecutionPlanSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionPlanSummary struct{}"
	}

	return strings.Join([]string{"ExecutionPlanSummary", string(data)}, " ")
}
