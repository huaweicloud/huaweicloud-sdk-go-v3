package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoDto struct {

	// 在标签请求参数中使用以获取输出的下一部分。重复此操作，直到响应元素返回null。如果存在，则表示可用的输出比当前响应中包含的输出多。
	NextMarker string `json:"next_marker"`

	// 本页返回条目数量。
	CurrentCount *int32 `json:"current_count,omitempty"`
}

func (o PageInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoDto struct{}"
	}

	return strings.Join([]string{"PageInfoDto", string(data)}, " ")
}
