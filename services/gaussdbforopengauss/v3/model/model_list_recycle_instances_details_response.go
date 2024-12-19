package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecycleInstancesDetailsResponse Response Object
type ListRecycleInstancesDetailsResponse struct {

	// 回收站所有引擎实例总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 回收站所有引擎实例信息。
	Instances      *[]RecycleInstancesDetailResultV1 `json:"instances,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListRecycleInstancesDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleInstancesDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListRecycleInstancesDetailsResponse", string(data)}, " ")
}
