package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DownloadSlowLogFileItem struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 文件名。
	FileName *string `json:"file_name,omitempty"`

	// 状态。  取值范围:   - SUCCESS：表示下载链接已经生成完成。   - EXPORTING：表示下载链接正在生成中。   - FAILED： 表示下载链接生成失败。
	Status *string `json:"status,omitempty"`

	// 文件大小，单位：KB。
	FileSize *string `json:"file_size,omitempty"`

	// 下载链接。链接有效时间为5分钟。
	FileLink *string `json:"file_link,omitempty"`

	// 创建时间。
	CreateAt *int64 `json:"create_at,omitempty"`

	// 更新时间。
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o DownloadSlowLogFileItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSlowLogFileItem struct{}"
	}

	return strings.Join([]string{"DownloadSlowLogFileItem", string(data)}, " ")
}
