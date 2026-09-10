package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduledEventRequestBody 修改计划事件预约时间请求体。
type UpdateScheduledEventRequestBody struct {

	// 计划执行开始时间，预约的时间需要比当前时间多5分钟以上，传空字符串表示立即执行。
	NotBefore string `json:"not_before"`
}

func (o UpdateScheduledEventRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduledEventRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateScheduledEventRequestBody", string(data)}, " ")
}
