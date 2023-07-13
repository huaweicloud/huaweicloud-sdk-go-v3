package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWholeImageResponse Response Object
type CreateWholeImageResponse struct {

	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWholeImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWholeImageResponse struct{}"
	}

	return strings.Join([]string{"CreateWholeImageResponse", string(data)}, " ")
}
