package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlinkJarRequest Request Object
type UpdateFlinkJarRequest struct {

	// 作业ID。
	JobId int64 `json:"job_id"`

	Body *UpdateFlinkJarRequestBody `json:"body,omitempty"`
}

func (o UpdateFlinkJarRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkJarRequest struct{}"
	}

	return strings.Join([]string{"UpdateFlinkJarRequest", string(data)}, " ")
}
