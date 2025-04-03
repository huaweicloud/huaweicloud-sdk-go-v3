package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SortDto struct {

	// 一个布尔值，表示服务提供商是否支持这种操作
	Supported *bool `json:"supported,omitempty"`
}

func (o SortDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SortDto struct{}"
	}

	return strings.Join([]string{"SortDto", string(data)}, " ")
}
