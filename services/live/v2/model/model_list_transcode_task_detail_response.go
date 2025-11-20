package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTranscodeTaskDetailResponse Response Object
type ListTranscodeTaskDetailResponse struct {

	// 推流域名
	Domain *string `json:"domain,omitempty"`

	// 采样开始时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	StartTime *string `json:"start_time,omitempty"`

	// 采样结束时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	EndTime *string `json:"end_time,omitempty"`

	// 流粒度转码明细
	TranscodeDetailList *[]TranscodeDetailInfo `json:"transcode_detail_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTranscodeTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTranscodeTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ListTranscodeTaskDetailResponse", string(data)}, " ")
}
