package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChartConfig struct {
	// 是否开启排序

	CanSort bool `json:"can_sort"`
	// 是否开启搜索

	CanSearch bool `json:"can_search"`
	// 每页显示数量

	PageSize int32 `json:"page_size"`
}

func (o ChartConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChartConfig struct{}"
	}

	return strings.Join([]string{"ChartConfig", string(data)}, " ")
}
