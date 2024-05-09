package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMolBatchDownloadTaskResponse Response Object
type CreateMolBatchDownloadTaskResponse struct {

	// 任务ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMolBatchDownloadTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMolBatchDownloadTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateMolBatchDownloadTaskResponse", string(data)}, " ")
}
