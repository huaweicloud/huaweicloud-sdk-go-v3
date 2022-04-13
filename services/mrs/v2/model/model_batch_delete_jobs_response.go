package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
