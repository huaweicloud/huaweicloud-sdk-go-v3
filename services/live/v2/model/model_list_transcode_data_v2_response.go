/*
 * LiveAPI
 *
 * 直播服务源站所有接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListTranscodeDataV2Response struct {
	// 采样数据列表。
	TranscodeDataList *[]TranscodeData `json:"transcode_data_list,omitempty"`
	// 指定时间区间内各转码规格转码时长总和。
	SummaryList *[]TranscodeSummary `json:"summary_list,omitempty"`
	XRequestId  *string             `json:"X-request-id,omitempty"`
}

func (o ListTranscodeDataV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListTranscodeDataV2Response", string(data)}, " ")
}
