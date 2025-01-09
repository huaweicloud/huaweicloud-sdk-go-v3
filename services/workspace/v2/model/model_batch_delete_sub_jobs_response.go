package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSubJobsResponse Response Object
type BatchDeleteSubJobsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSubJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubJobsResponse", string(data)}, " ")
}
