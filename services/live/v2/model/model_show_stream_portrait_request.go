package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowStreamPortraitRequest struct {
	// 播放域名。

	PlayDomain string `json:"play_domain"`
	// 流名。

	Stream *string `json:"stream,omitempty"`
	// 统计日期，日期格式按照ISO8601表示法，格式：YYYYMMDD，如20200904。可以查询过去31天的数据（不含当天）。

	Time string `json:"time"`
}

func (o ShowStreamPortraitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStreamPortraitRequest struct{}"
	}

	return strings.Join([]string{"ShowStreamPortraitRequest", string(data)}, " ")
}
