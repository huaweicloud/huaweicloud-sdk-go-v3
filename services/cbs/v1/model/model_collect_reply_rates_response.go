package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectReplyRatesResponse Response Object
type CollectReplyRatesResponse struct {

	// 统计周期目前支持year、month、week、day。 调用失败时无此字段。
	Interval *string `json:"interval,omitempty"`

	// 所在时区，例如：中国东八区为\"+08:00\"；美国西五区为\"-05:00\";默认为\"UTC\"。 调用失败时无此字段。
	TimeZone *string `json:"time_zone,omitempty"`

	Total *ReplyRatesTotal `json:"total,omitempty"`

	// 会话间隔统计数据。
	Intervals *[]ReplyRatesIntervals `json:"intervals,omitempty"`

	// 统计开始的utc时间。
	Startutc *int64 `json:"startutc,omitempty"`

	// 统计结束的utc时间。
	Endutc         *int64 `json:"endutc,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CollectReplyRatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectReplyRatesResponse struct{}"
	}

	return strings.Join([]string{"CollectReplyRatesResponse", string(data)}, " ")
}
