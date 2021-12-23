package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSingleStreamFramerateRequest struct {
	// 推流域名。

	Domain string `json:"domain"`
	// App名。

	App string `json:"app"`
	// 流名。

	Stream string `json:"stream"`
	// 起始时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度1天，最大查询周期1个月。  若参数为空，默认查询最近1小时数据。

	StartTime *string `json:"start_time,omitempty"`
	// 结束时间。日期格式按照ISO8601表示法，并使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。  若参数为空，默认为当前时间。结束时间需大于起始时间。

	EndTime *string `json:"end_time,omitempty"`
}

func (o ListSingleStreamFramerateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSingleStreamFramerateRequest struct{}"
	}

	return strings.Join([]string{"ListSingleStreamFramerateRequest", string(data)}, " ")
}
