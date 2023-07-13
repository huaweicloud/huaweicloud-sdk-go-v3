package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScrumProjectStatusesRequest Request Object
type ListScrumProjectStatusesRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 一次返回的数据,最小1,最大100
	Limit *int32 `json:"limit,omitempty"`

	// 自定义字段支持的工作项类型 2任务/Task,3缺陷/Bug,5Epic,6Feature,7Story
	TrackerId *int32 `json:"tracker_id,omitempty"`
}

func (o ListScrumProjectStatusesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScrumProjectStatusesRequest struct{}"
	}

	return strings.Join([]string{"ListScrumProjectStatusesRequest", string(data)}, " ")
}
