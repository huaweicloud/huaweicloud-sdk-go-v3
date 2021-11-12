package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateImageResponse struct {
	// 异步任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageResponse struct{}"
	}

	return strings.Join([]string{"CreateImageResponse", string(data)}, " ")
}
