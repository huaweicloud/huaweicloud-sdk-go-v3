package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorGroupInfo struct {

	// 计算资源架构类型，目前分X86和ARM两种。
	GroupType string `json:"group_type"`

	// 规格详情。
	Flavors []Flavor `json:"flavors"`

	// 分页参数: 起始值。
	Offset int32 `json:"offset"`

	// 分页参数：每页多少条。
	Limit int32 `json:"limit"`

	// 计算类型规格总数。
	Total int32 `json:"total"`
}

func (o FlavorGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorGroupInfo struct{}"
	}

	return strings.Join([]string{"FlavorGroupInfo", string(data)}, " ")
}
