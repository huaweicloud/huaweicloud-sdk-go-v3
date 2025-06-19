package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearRecyclingJobsResponse Response Object
type ClearRecyclingJobsResponse struct {

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ClearRecyclingJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearRecyclingJobsResponse struct{}"
	}

	return strings.Join([]string{"ClearRecyclingJobsResponse", string(data)}, " ")
}
