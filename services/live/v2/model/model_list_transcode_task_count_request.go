package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTranscodeTaskCountRequest struct {
	// 推流域名

	PublishDomain string `json:"publish_domain"`
	// 应用名称，若查询结果为空，表示该应用下没有转码任务。

	App *string `json:"app,omitempty"`
	// 查询起始时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度7天，最大查询周期90天。  若参数为空，默认查询7天数据。

	StartTime *string `json:"start_time,omitempty"`
	// 查询结束时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度7天，最大查询周期90天。  结束时间需大于起始时间。  若参数为空，默认为当前时间。

	EndTime *string `json:"end_time,omitempty"`
}

func (o ListTranscodeTaskCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTranscodeTaskCountRequest struct{}"
	}

	return strings.Join([]string{"ListTranscodeTaskCountRequest", string(data)}, " ")
}
