package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAsyncTtsJobResponse Response Object
type CreateAsyncTtsJobResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAsyncTtsJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAsyncTtsJobResponse struct{}"
	}

	return strings.Join([]string{"CreateAsyncTtsJobResponse", string(data)}, " ")
}
