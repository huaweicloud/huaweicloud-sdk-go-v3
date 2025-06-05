package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatisticDto 资源统计模型
type StatisticDto struct {

	// 统计资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 统计资源值
	Count *int64 `json:"count,omitempty"`
}

func (o StatisticDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticDto struct{}"
	}

	return strings.Join([]string{"StatisticDto", string(data)}, " ")
}
