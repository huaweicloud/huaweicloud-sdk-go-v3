package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTranscodeDataResponse struct {
	// 采样数据列表。

	TranscodeDataList *[]TranscodeData `json:"transcode_data_list,omitempty"`
	// 指定时间区间内各转码规格转码时长总和。

	SummaryList *[]TranscodeSummary `json:"summary_list,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTranscodeDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTranscodeDataResponse struct{}"
	}

	return strings.Join([]string{"ListTranscodeDataResponse", string(data)}, " ")
}
