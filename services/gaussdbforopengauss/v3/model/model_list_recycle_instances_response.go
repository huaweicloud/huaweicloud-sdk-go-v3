package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRecycleInstancesResponse struct {

	// 回收站所有引擎实例总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 回收站所有引擎实例信息。
	Instances      *[]RecycleInstancesDetailResult `json:"instances,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListRecycleInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListRecycleInstancesResponse", string(data)}, " ")
}
