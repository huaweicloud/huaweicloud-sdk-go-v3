package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryJobsResponse Response Object
type RetryJobsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RetryJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryJobsResponse struct{}"
	}

	return strings.Join([]string{"RetryJobsResponse", string(data)}, " ")
}
