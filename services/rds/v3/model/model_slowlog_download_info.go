package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowlogDownloadInfo struct {

	// 任务ID
	WorkflowId string `json:"workflow_id" xml:"workflow_id"`

	// 生成的下载文件名
	FileName string `json:"file_name" xml:"file_name"`

	// 生成链接的生成状态
	Status string `json:"status" xml:"status"`

	// 文件大小
	FileSize string `json:"file_size" xml:"file_size"`

	// 下载链接
	FileLink string `json:"file_link" xml:"file_link"`

	// 生成时间
	CreateAt int64 `json:"create_at" xml:"create_at"`

	// 更新时间
	UpdateAt int64 `json:"update_at" xml:"update_at"`
}

func (o SlowlogDownloadInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowlogDownloadInfo struct{}"
	}

	return strings.Join([]string{"SlowlogDownloadInfo", string(data)}, " ")
}
