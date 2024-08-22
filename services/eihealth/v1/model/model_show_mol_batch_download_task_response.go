package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMolBatchDownloadTaskResponse Response Object
type ShowMolBatchDownloadTaskResponse struct {

	// 任务状态：WAITING、RUNNING、FINISHED、CANCELLED、ABNORMAL、FAILED
	Status *string `json:"status,omitempty"`

	// 下载文件名
	Filename *string `json:"filename,omitempty"`

	// 下载路径
	OutDir *string `json:"out_dir,omitempty"`

	Progress       *TaskProgress `json:"progress,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowMolBatchDownloadTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMolBatchDownloadTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowMolBatchDownloadTaskResponse", string(data)}, " ")
}
