package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveEventReportRequest Request Object
type LiveEventReportRequest struct {

	// 剧本ID。
	RoomId string `json:"room_id"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 鉴权Key。通过HmacSHA256生成的鉴权key
	AuthKey *string `json:"auth_key,omitempty"`

	// **参数解释**： 鉴权key过期时间。从1970年1月1日（UTC/GMT的午夜）开始所经过的豪秒数。
	ExpiresTime *int64 `json:"expires_time,omitempty"`

	// 是否刷新URL
	RefreshUrl *bool `json:"refresh_url,omitempty"`

	Body *ReportLiveEventReq `json:"body,omitempty"`
}

func (o LiveEventReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveEventReportRequest struct{}"
	}

	return strings.Join([]string{"LiveEventReportRequest", string(data)}, " ")
}
