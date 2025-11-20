package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTranscodeTaskDetailRequest Request Object
type ListTranscodeTaskDetailRequest struct {

	// 推流域名。
	Domain string `json:"domain"`

	// 流名列表，以逗号分隔，最多支持查询100个流名。 如果不传入流名，则查询域名下所有转码流的数据。
	StreamNameList *[]string `json:"stream_name_list,omitempty"`

	// 起始时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度1天，最大查询周期14天。  若参数为空，默认查询1天数据。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度1天，最大查询周期14天。  若参数为空，默认为当前时间。结束时间需大于起始时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListTranscodeTaskDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTranscodeTaskDetailRequest struct{}"
	}

	return strings.Join([]string{"ListTranscodeTaskDetailRequest", string(data)}, " ")
}
