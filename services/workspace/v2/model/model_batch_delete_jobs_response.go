package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteJobsResponse Response Object
type BatchDeleteJobsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobsResponse", string(data)}, " ")
}
