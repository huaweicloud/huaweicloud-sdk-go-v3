package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobHistoryParametersRequest Request Object
type ListJobHistoryParametersRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset 大于等于 0。默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。默认为10，取值范围【1-1000】
	Limit *int32 `json:"limit,omitempty"`

	// 开始时间，UTC时间，例如：2020-09-01T18:50:20Z
	BeginTime *string `json:"begin_time,omitempty"`

	// 结束时间，UTC时间，例如：2020-09-01T19:50:20Z
	EndTime *string `json:"end_time,omitempty"`

	// 参数名称。
	Name *string `json:"name,omitempty"`
}

func (o ListJobHistoryParametersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobHistoryParametersRequest struct{}"
	}

	return strings.Join([]string{"ListJobHistoryParametersRequest", string(data)}, " ")
}
