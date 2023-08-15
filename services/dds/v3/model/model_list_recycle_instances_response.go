package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecycleInstancesResponse Response Object
type ListRecycleInstancesResponse struct {

	// 总记录数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 实例信息
	Instances      *[]RecycleInstance `json:"instances,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListRecycleInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListRecycleInstancesResponse", string(data)}, " ")
}
