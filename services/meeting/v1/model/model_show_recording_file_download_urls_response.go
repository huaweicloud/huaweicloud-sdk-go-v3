package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRecordingFileDownloadUrlsResponse struct {

	// 录制文件下载URL
	RecordUrls     *[]RecordDownloadInfoBo `json:"recordUrls,omitempty" xml:"recordUrls"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowRecordingFileDownloadUrlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordingFileDownloadUrlsResponse struct{}"
	}

	return strings.Join([]string{"ShowRecordingFileDownloadUrlsResponse", string(data)}, " ")
}
