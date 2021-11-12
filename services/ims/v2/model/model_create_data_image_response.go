package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDataImageResponse struct {
	// 异步任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDataImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataImageResponse struct{}"
	}

	return strings.Join([]string{"CreateDataImageResponse", string(data)}, " ")
}
