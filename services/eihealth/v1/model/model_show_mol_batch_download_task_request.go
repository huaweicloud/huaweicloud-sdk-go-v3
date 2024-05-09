package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMolBatchDownloadTaskRequest Request Object
type ShowMolBatchDownloadTaskRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 下载任务ID。
	TaskId string `json:"task_id"`
}

func (o ShowMolBatchDownloadTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMolBatchDownloadTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowMolBatchDownloadTaskRequest", string(data)}, " ")
}
