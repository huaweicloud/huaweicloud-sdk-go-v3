package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FilterDto struct {

	// 一个布尔值，表示服务提供商是否支持这种操作
	Supported *bool `json:"supported,omitempty"`

	// 最大结果数
	MaxResults *int32 `json:"maxResults,omitempty"`
}

func (o FilterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilterDto struct{}"
	}

	return strings.Join([]string{"FilterDto", string(data)}, " ")
}
