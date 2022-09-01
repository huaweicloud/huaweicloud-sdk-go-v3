package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRtcHistoryUsageRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty" xml:"X-Sdk-Date"`

	// 使用AK/SK方式认证时必选，携带项目ID信息，与路径参数中的项目ID相同。
	XProjectId *string `json:"X-Project-Id,omitempty" xml:"X-Project-Id"`

	// 应用标识
	App string `json:"app" xml:"app"`

	// 查询的数据类型 - CommunicationDuration: 音视频通话时长; - TranscodeDuration：转码时长; - RecordDuration：录制时长;
	Metric string `json:"metric" xml:"metric"`

	// 查询起始时间。UTC时间，格式：YYYY-MM-DD，如2020-04-23。
	StartDate string `json:"start_date" xml:"start_date"`

	// 查询结束时间。UTC时间，格式：YYYY-MM-DD，如2020-04-23。
	EndDate string `json:"end_date" xml:"end_date"`
}

func (o ListRtcHistoryUsageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcHistoryUsageRequest struct{}"
	}

	return strings.Join([]string{"ListRtcHistoryUsageRequest", string(data)}, " ")
}
