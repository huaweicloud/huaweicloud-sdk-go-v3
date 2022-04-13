package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 延长会议消息体。
type RestProlongDurReqBody struct {
	// - 0: 手动延长。 - 1: 自动延长（未携带延长时间时，默认每次延长15分钟）。

	Auto int32 `json:"auto"`
	// 延长时间，单位为分钟。 默认值：15

	Duration *int32 `json:"duration,omitempty"`
}

func (o RestProlongDurReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestProlongDurReqBody struct{}"
	}

	return strings.Join([]string{"RestProlongDurReqBody", string(data)}, " ")
}
