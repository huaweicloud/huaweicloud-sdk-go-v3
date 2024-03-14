package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoDto struct {

	// 如果存在，则表示可用的输出比当前响应中包含的输出多。在对操作的后续调用中，在标签请求参数中使用此值，以获取输出的下一部分。您应该重复此操作，直到next_marker响应元素返回为null
	NextMarker string `json:"next_marker"`

	// 本页返回条目数量
	CurrentCount *int32 `json:"current_count,omitempty"`
}

func (o PageInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoDto struct{}"
	}

	return strings.Join([]string{"PageInfoDto", string(data)}, " ")
}
