package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRecordingFileDownloadUrlsRequest struct {

	// 会议的ConfUUID(通过查询录制文件列表获取)。
	ConfUUID string `json:"confUUID" xml:"confUUID"`

	// 数据偏移记录。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 指定返回的记录数，最大500条。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 标识是否为第三方portal过来的请求。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty" xml:"X-Authorization-Type"`

	// 用于区分到哪个HCSO站点鉴权。
	XSiteId *string `json:"X-Site-Id,omitempty" xml:"X-Site-Id"`
}

func (o ShowRecordingFileDownloadUrlsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordingFileDownloadUrlsRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordingFileDownloadUrlsRequest", string(data)}, " ")
}
