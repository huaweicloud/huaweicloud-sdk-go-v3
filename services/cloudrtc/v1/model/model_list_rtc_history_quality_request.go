package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRtcHistoryQualityRequest Request Object
type ListRtcHistoryQualityRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息，与路径参数中的项目ID相同。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 应用标识
	App string `json:"app"`

	// 查询的数据类型 - JoinSuccessRate：加入房间成功率 - JoinSuccess5SecsRate：5秒加入成功率 - VideoFreezeRate：视频卡顿率 - AudioFreezeRate：音频卡顿率 - FirstVideoRecvTime：首帧视频接收耗时 - FirstAudioRecvTime：首帧音频接收耗时 - PullStreamSuccessRate：拉流成功率 - PushStreamSuccessRate：推流成功率 - VideoUpstreamExcellentTransRate：客户端视频上行优质传输率 - AudioUpstreamExcellentTransRate：客户端音频上行优质传输率 - VideoExcellentTransRate：端到端视频优质传输率 - AudioExcellentTransRate：端到端音频优质传输率 - VideoTransDelay：端到端视频网络时，单位为毫秒，取当天所有用户网络延迟的中位数 - AudioTransDelay：端到端音频网络时延，单位为毫秒，取当天所有用户网络延迟的中位数
	Metric []string `json:"metric"`

	// 查询起始时间。UTC时间，格式：YYYY-MM-DD，如2020-04-23，不写默认读取过去1天数据数据。
	StartDate *string `json:"start_date,omitempty"`

	// 查询结束时间。UTC时间，格式：YYYY-MM-DD，如2020-04-23
	EndDate *string `json:"end_date,omitempty"`
}

func (o ListRtcHistoryQualityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcHistoryQualityRequest struct{}"
	}

	return strings.Join([]string{"ListRtcHistoryQualityRequest", string(data)}, " ")
}
