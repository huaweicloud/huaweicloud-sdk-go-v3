package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHistoryStreamsRequest struct {

	// 推流域名。
	Domain string `json:"domain" xml:"domain"`

	// 应用名称。
	App *string `json:"app,omitempty" xml:"app"`

	// 流名称。
	Stream *string `json:"stream,omitempty" xml:"stream"`

	// 起始时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度1天。  若参数为空，默认查询1天数据。
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 结束时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。  若参数为空，默认为当前时间，最大查询跨度1天。结束时间需大于起始时间。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 分页编号，默认为0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页记录数。  取值范围：[1,100]  默认值：10。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListHistoryStreamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryStreamsRequest struct{}"
	}

	return strings.Join([]string{"ListHistoryStreamsRequest", string(data)}, " ")
}
