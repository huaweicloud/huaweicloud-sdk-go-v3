package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群维护时间窗信息
type MaintainWindow struct {

	// 每周的维护时间，以天为粒度，取值如下：  - Mon：星期一 - Tue：星期二 - Wed：星期三 - Thu：星期四 - Fri：星期五 - Sat：星期六 - Sun：星期日
	Day *string `json:"day,omitempty" xml:"day"`

	// 维护开始时间，显示格式为 HH：mm，时区为GMT+0。
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 维护结束时间，显示格式为 HH：mm, 时区为GMT+0。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`
}

func (o MaintainWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaintainWindow struct{}"
	}

	return strings.Join([]string{"MaintainWindow", string(data)}, " ")
}
