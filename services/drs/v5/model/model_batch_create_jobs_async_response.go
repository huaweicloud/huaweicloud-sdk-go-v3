package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateJobsAsyncResponse struct {
	Job            *AsyncCreateJobResp `json:"job,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o BatchCreateJobsAsyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateJobsAsyncResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateJobsAsyncResponse", string(data)}, " ")
}
