package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHistoryStreamsRequest struct {
	// 推流域名。

	Domain string `json:"domain"`
	// 应用名称。

	App *string `json:"app,omitempty"`
	// 流名称。

	Stream *string `json:"stream,omitempty"`
	// 起始时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度1天，最大查询周期1个月。  若参数为空，默认查询1天数据。

	StartTime *string `json:"start_time,omitempty"`
	// 结束时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。  若参数为空，默认为当前时间，最大查询跨度1天，最大查询周期1个月。结束时间需大于起始时间。

	EndTime *string `json:"end_time,omitempty"`
	// 分页编号，默认为0

	Offset *int32 `json:"offset,omitempty"`
	// 每页记录数。  取值范围：[1,100]  默认值：10。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListHistoryStreamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryStreamsRequest struct{}"
	}

	return strings.Join([]string{"ListHistoryStreamsRequest", string(data)}, " ")
}
