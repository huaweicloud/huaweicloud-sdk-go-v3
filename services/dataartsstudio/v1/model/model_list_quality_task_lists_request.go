package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQualityTaskListsRequest Request Object
type ListQualityTaskListsRequest struct {

	// start number
	Start *int64 `json:"start,omitempty"`

	// page size
	PageSize *int64 `json:"page_size,omitempty"`

	// 分页查询偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// category id
	CategoryId *int64 `json:"category_id,omitempty"`

	// rule name
	RuleName *string `json:"rule_name,omitempty"`

	// schedule status
	ScheduleStatus *int32 `json:"schedule_status,omitempty"`

	// schedule period
	SchedulePeriod *int32 `json:"schedule_period,omitempty"`

	// 开始时间(搜索)
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间(搜索)
	EndTime *string `json:"end_time,omitempty"`

	// 最近运行结果 0：运行中 1：异常 2：告警 3：正常
	ResultStatus *int32 `json:"result_status,omitempty"`

	// 排序字段
	Sort *string `json:"sort,omitempty"`

	// 排序方式
	Order *string `json:"order,omitempty"`
}

func (o ListQualityTaskListsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQualityTaskListsRequest struct{}"
	}

	return strings.Join([]string{"ListQualityTaskListsRequest", string(data)}, " ")
}
