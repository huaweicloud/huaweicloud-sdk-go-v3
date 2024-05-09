package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMolBatchDownloadTaskRequest Request Object
type CreateMolBatchDownloadTaskRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateMolBatchDownloadTaskReq `json:"body,omitempty"`
}

func (o CreateMolBatchDownloadTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMolBatchDownloadTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateMolBatchDownloadTaskRequest", string(data)}, " ")
}
