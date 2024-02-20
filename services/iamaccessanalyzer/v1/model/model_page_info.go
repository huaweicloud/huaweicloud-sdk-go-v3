package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageInfo 页面的信息。
type PageInfo struct {

	// 当前页中的项数。
	CurrentCount int32 `json:"current_count"`

	// 如果存在更多可用的输出，那么该值表示可用输出比当前响应中包含的更多。在后续调用此操作时，您可以在标记请求参数中使用此值，以获取输出的下一部分。您应该重复这个过程，直到next_marker返回为null。
	NextMarker *string `json:"next_marker,omitempty"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
