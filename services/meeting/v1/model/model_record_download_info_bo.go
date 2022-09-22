package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会议录制文件下载链接信息。
type RecordDownloadInfoBo struct {

	// 会议UUID。
	ConfUuid *string `json:"confUuid,omitempty"`

	// 下载链接信息。
	Urls *[]RecordDownloadUrlDo `json:"urls,omitempty"`
}

func (o RecordDownloadInfoBo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordDownloadInfoBo struct{}"
	}

	return strings.Join([]string{"RecordDownloadInfoBo", string(data)}, " ")
}
