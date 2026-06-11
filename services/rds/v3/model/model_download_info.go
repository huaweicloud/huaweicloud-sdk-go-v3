package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadInfo 下载信息
type DownloadInfo struct {

	// 参数解释： 任务ID。 取值范围： 不涉及。
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 参数解释： 生成的下载文件名。 取值范围： 不涉及。
	FileName *string `json:"file_name,omitempty"`

	// 参数解释： 当前链接的生成状态。 取值范围： 不涉及。
	Status *string `json:"status,omitempty"`

	// 参数解释： 文件大小。单位Byte 取值范围： 不涉及。
	FileSize *string `json:"file_size,omitempty"`

	// 参数解释： 下载链接。 取值范围： 不涉及。
	FileLink *string `json:"file_link,omitempty"`

	// 下载链接过期时间
	FileLinkExpirationTime *int64 `json:"file_link_expiration_time,omitempty"`

	// 参数解释： 开始时间。 格式为UTC时间戳。 取值范围： 不涉及。
	StartTime *int64 `json:"start_time,omitempty"`

	// 参数解释： 结束时间。 格式为UTC时间戳。 取值范围： 不涉及。
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o DownloadInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadInfo struct{}"
	}

	return strings.Join([]string{"DownloadInfo", string(data)}, " ")
}
