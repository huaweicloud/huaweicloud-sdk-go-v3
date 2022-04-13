package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRecordingFileDownloadUrlsRequest struct {
	// 会议的ConfUUID(通过查询录制文件列表获取)。

	ConfUUID string `json:"confUUID"`
	// 数据偏移记录。

	Offset *int32 `json:"offset,omitempty"`
	// 指定返回的记录数，最大500条。

	Limit *int32 `json:"limit,omitempty"`
	// 标识是否为第三方portal过来的请求。

	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`
	// 用于区分到哪个HCSO站点鉴权。

	XSiteId *string `json:"X-Site-Id,omitempty"`
}

func (o ShowRecordingFileDownloadUrlsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordingFileDownloadUrlsRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordingFileDownloadUrlsRequest", string(data)}, " ")
}
