package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimMsgAppRequest Request Object
type ListAimMsgAppRequest struct {

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 应用状态。
	Status *string `json:"status,omitempty"`

	// 创建时间筛选-开始时间。格式为：2019-10-12T07:20:50Z。
	BeginTime *string `json:"begin_time,omitempty"`

	// 创建时间筛选-结束时间。格式为：2019-10-12T07:20:50Z。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListAimMsgAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimMsgAppRequest struct{}"
	}

	return strings.Join([]string{"ListAimMsgAppRequest", string(data)}, " ")
}
