package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountSubJobsResponse Response Object
type CountSubJobsResponse struct {

	// 总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountSubJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountSubJobsResponse struct{}"
	}

	return strings.Join([]string{"CountSubJobsResponse", string(data)}, " ")
}
