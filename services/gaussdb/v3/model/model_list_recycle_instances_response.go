package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecycleInstancesResponse Response Object
type ListRecycleInstancesResponse struct {

	// 数据总数
	TotalCount *string `json:"total_count,omitempty"`

	// 回收站实例信息
	Instances      *[]RecycleInstanceV3 `json:"instances,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListRecycleInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListRecycleInstancesResponse", string(data)}, " ")
}
