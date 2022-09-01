package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowProjectWorkHoursRequestBody struct {

	// 查询的用户id列表
	UserIds *[]string `json:"user_ids,omitempty" xml:"user_ids"`

	// 工时类型，以逗号分隔,21:研发设计,22:后端开发,23:前端开发(Web),24:前端开发(小程序),25:前端开发(App),26:测试验证,27:缺陷修复,28:UI设计,29:会议,30:公共事务,31:培训,32:研究,33:其它,34:调休请假
	WorkHoursTypes *string `json:"work_hours_types,omitempty" xml:"work_hours_types"`

	// 工时日期，以逗号分隔，年-月-日
	WorkHoursDates *string `json:"work_hours_dates,omitempty" xml:"work_hours_dates"`

	// 工时开始日期，年-月-日
	BeginTime *string `json:"begin_time,omitempty" xml:"begin_time"`

	// 工时结束日期，年-月-日
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 偏移量,offset是limit的整数倍，limit=10,offset=0,10,20...
	Offset int32 `json:"offset" xml:"offset"`

	// 每页显示数量，每页最多显示100条
	Limit int32 `json:"limit" xml:"limit"`
}

func (o ShowProjectWorkHoursRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectWorkHoursRequestBody struct{}"
	}

	return strings.Join([]string{"ShowProjectWorkHoursRequestBody", string(data)}, " ")
}
