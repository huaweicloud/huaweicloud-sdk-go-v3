package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCarouselTaskDetailRequest Request Object
type ListCarouselTaskDetailRequest struct {

	// 轮播任务id。
	CarouselTaskId string `json:"carousel_task_id"`

	// 起始时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度3小时，最大查询周期7天。  若参数为空，默认查询最近1小时数据。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度3小时，最大查询周期7天。  若参数为空，默认为当前时间。结束时间需大于起始时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListCarouselTaskDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCarouselTaskDetailRequest struct{}"
	}

	return strings.Join([]string{"ListCarouselTaskDetailRequest", string(data)}, " ")
}
