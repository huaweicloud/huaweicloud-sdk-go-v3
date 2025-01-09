package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteImageSubJobsResponse Response Object
type BatchDeleteImageSubJobsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteImageSubJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteImageSubJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteImageSubJobsResponse", string(data)}, " ")
}
