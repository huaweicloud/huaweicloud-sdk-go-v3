package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NextMarker 如果存在更多可用的输出，那么该值表示可用输出比当前响应中包含的更多。在后续调用此操作时，您可以在标记请求参数中使用此值，以获取输出的下一部分。您应该重复这个过程，直到next_marker返回为null。
type NextMarker struct {
}

func (o NextMarker) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NextMarker struct{}"
	}

	return strings.Join([]string{"NextMarker", string(data)}, " ")
}
