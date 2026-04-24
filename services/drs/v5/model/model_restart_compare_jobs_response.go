package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartCompareJobsResponse Response Object
type RestartCompareJobsResponse struct {

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartCompareJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartCompareJobsResponse struct{}"
	}

	return strings.Join([]string{"RestartCompareJobsResponse", string(data)}, " ")
}
