package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRunsResponse struct {
	// 作业总个数。

	Count *int64 `json:"count,omitempty"`
	// 作业运行列表。

	Runs           *[]Run `json:"runs,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRunsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRunsResponse struct{}"
	}

	return strings.Join([]string{"ListRunsResponse", string(data)}, " ")
}
