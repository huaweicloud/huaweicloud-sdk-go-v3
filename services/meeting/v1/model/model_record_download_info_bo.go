package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 单会议录制文件下载链接信息（包含多个录制文件）
type RecordDownloadInfoBo struct {
	// 会议UUID

	ConfUuid *string `json:"confUuid,omitempty"`
	// 录制文件下载URL

	Urls *[]RecordDownloadUrlDo `json:"urls,omitempty"`
}

func (o RecordDownloadInfoBo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordDownloadInfoBo struct{}"
	}

	return strings.Join([]string{"RecordDownloadInfoBo", string(data)}, " ")
}
