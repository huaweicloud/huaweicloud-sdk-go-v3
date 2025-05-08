package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUploadByUrlResponse Response Object
type CreateUploadByUrlResponse struct {

	// 拉取任务id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateUploadByUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUploadByUrlResponse struct{}"
	}

	return strings.Join([]string{"CreateUploadByUrlResponse", string(data)}, " ")
}
