package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAppSubJobsResponse Response Object
type BatchDeleteAppSubJobsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteAppSubJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAppSubJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAppSubJobsResponse", string(data)}, " ")
}
