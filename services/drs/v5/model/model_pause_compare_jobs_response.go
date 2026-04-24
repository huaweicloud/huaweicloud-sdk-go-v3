package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PauseCompareJobsResponse Response Object
type PauseCompareJobsResponse struct {

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PauseCompareJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseCompareJobsResponse struct{}"
	}

	return strings.Join([]string{"PauseCompareJobsResponse", string(data)}, " ")
}
