package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowIterationV4Response struct {

	// 迭代结束时间，年-月-日
	BeginTime *string `json:"begin_time,omitempty" xml:"begin_time"`

	// 燃尽图
	Charts *[]Chart `json:"charts,omitempty" xml:"charts"`

	// 已关闭的工单数
	ClosedTotal *int32 `json:"closed_total,omitempty" xml:"closed_total"`

	// 迭代创建时间
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 迭代开始时间，年-月-日
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 是否有task
	HaveTask *bool `json:"have_task,omitempty" xml:"have_task"`

	// 迭代id
	IterationId *int32 `json:"iteration_id,omitempty" xml:"iteration_id"`

	// 迭代标题
	Name *string `json:"name,omitempty" xml:"name"`

	// 开启的工单数
	OpenedTotal *int32 `json:"opened_total,omitempty" xml:"opened_total"`

	// 工作进展
	Progress *string `json:"progress,omitempty" xml:"progress"`

	// 工单总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 迭代更新时间
	UpdatedTime *string `json:"updated_time,omitempty" xml:"updated_time"`

	// 迭代的状态,0 未开始，1 进行中，2 结束
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowIterationV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIterationV4Response struct{}"
	}

	return strings.Join([]string{"ShowIterationV4Response", string(data)}, " ")
}
