package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSinkTaskDetailResponse struct {

	// 转储任务名称。
	TaskName *string `json:"task_name,omitempty" xml:"task_name"`

	// 转储任务类型。
	DestinationType *string `json:"destination_type,omitempty" xml:"destination_type"`

	// 转储任务创建时间戳。
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 转储任务状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 返回任务转存的topics列表或者正则表达式。
	Topics *string `json:"topics,omitempty" xml:"topics"`

	ObsDestinationDescriptor *ShowSinkTaskDetailRespObsDestinationDescriptor `json:"obs_destination_descriptor,omitempty" xml:"obs_destination_descriptor"`

	// topic信息。
	TopicsInfo     *[]ShowSinkTaskDetailRespTopicsInfo `json:"topics_info,omitempty" xml:"topics_info"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ShowSinkTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSinkTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowSinkTaskDetailResponse", string(data)}, " ")
}
