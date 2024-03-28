package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlinkJarJobRequest Request Object
type UpdateFlinkJarJobRequest struct {

	// 作业ID。
	JobId int64 `json:"job_id"`

	Body *UpdateFlinkJarJobRequestBody `json:"body,omitempty"`
}

func (o UpdateFlinkJarJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkJarJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateFlinkJarJobRequest", string(data)}, " ")
}
