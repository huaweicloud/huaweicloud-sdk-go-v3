package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCenterTaskRequest Request Object
type ListCenterTaskRequest struct {

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示条数，最小值为1，最大值为1000，若不设置该参数，则为10。
	Limit *int32 `json:"limit,omitempty"`

	// 查询开始时间，时间为UTC时间。格式：yyyyMMddHHmmss，如：20200609160000。
	StartTime *string `json:"start_time,omitempty"`

	// 查询结束时间，时间为UTC时间。格式：yyyyMMddHHmmss，如：20230612155959。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListCenterTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCenterTaskRequest struct{}"
	}

	return strings.Join([]string{"ListCenterTaskRequest", string(data)}, " ")
}
