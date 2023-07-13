package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CycleParams 周期会议的参数。
type CycleParams struct {

	// 周期会议的开始日期，格式：YYYY-MM-DD。 开始日期不能早于当前日期。 > 日期是timeZoneID指定的时区的日期，非UTC时间的日期。
	StartDate string `json:"startDate"`

	// 周期会议的结束日期，格式：YYYY-MM-DD。 开始日期和结束日期间的时间间隔最长不能超过1年。开始日期和结束日期之间最多允许50个子会议，若超过50个子会议，会自动调整结束日期。 > 日期是timeZoneID指定的时区的日期，非UTC时间的日期。
	EndDate string `json:"endDate"`

	// 周期类型。 - Day: 天 - Week: 星期 - Month: 月
	Cycle string `json:"cycle"`

	// 子会议间隔。 - cycle选择了Day，表示每几天召开一次，取值范围[1,15] - cycle选择了Week，表示每几周召开一次，取值范围[1,5] - cycle选择了Month，Interval表示隔几月，取值范围[1,3]
	Interval *int32 `json:"interval,omitempty"`

	// 周期内的会议召开点。仅当按周和月时有效。 - cycle选择了Week，point中填入了两个元素1和3，则表示每个周一和周三召开会议，0表示周日 - cycle选择了Month，point中填入了12和20则表示每个月的12号和20号召开会议，取值范围为[1,31]，若当月没有该值，则为月末
	Point *[]int32 `json:"point,omitempty"`

	// 提前通知天数。所有与会者在每个子会议开始前N天收到会议通知。取值范围[0,30]。 默认值是1。
	PreRemindDays int32 `json:"preRemindDays"`
}

func (o CycleParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CycleParams struct{}"
	}

	return strings.Join([]string{"CycleParams", string(data)}, " ")
}
