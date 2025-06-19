package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecyclingJobsResponse Response Object
type DeleteRecyclingJobsResponse struct {

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRecyclingJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecyclingJobsResponse struct{}"
	}

	return strings.Join([]string{"DeleteRecyclingJobsResponse", string(data)}, " ")
}
