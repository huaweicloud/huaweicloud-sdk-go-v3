package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountImageSubJobsResponse Response Object
type CountImageSubJobsResponse struct {

	// 总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountImageSubJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountImageSubJobsResponse struct{}"
	}

	return strings.Join([]string{"CountImageSubJobsResponse", string(data)}, " ")
}
