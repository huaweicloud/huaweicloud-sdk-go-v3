package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComputeFlavorGroupsInfo struct {
	// 计算资源架构类型，目前分X86和ARM两种。

	GroupType *string `json:"groupType,omitempty"`
	// 计算类型规格详情。

	ComputeFlavors *[]ComputeFlavors `json:"computeFlavors,omitempty"`
	// 分页参数: 起始值。

	Offset *int32 `json:"offset,omitempty"`
	// 分页参数：每页多少条。

	Limit *int32 `json:"limit,omitempty"`
	// 计算类型规格总数。

	Total *int32 `json:"total,omitempty"`
}

func (o ComputeFlavorGroupsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeFlavorGroupsInfo struct{}"
	}

	return strings.Join([]string{"ComputeFlavorGroupsInfo", string(data)}, " ")
}
