package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadHistorySessionAnalyaseBody 请求体
type UploadHistorySessionAnalyaseBody struct {

	// 类型。 session    表示下载历史会话 wait_event    表示下载历史等待事件 top_sql    表示下载历史top sql
	Type string `json:"type"`

	// 参数解释： 开始时间。 格式为UTC时间戳。 取值范围： 不涉及。 默认取值： 不涉及。
	StartTime int64 `json:"start_time"`

	// 参数解释： 结束时间。 格式为UTC时间戳。 取值范围： 不涉及。 默认取值： 不涉及。
	EndTime int64 `json:"end_time"`
}

func (o UploadHistorySessionAnalyaseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadHistorySessionAnalyaseBody struct{}"
	}

	return strings.Join([]string{"UploadHistorySessionAnalyaseBody", string(data)}, " ")
}
