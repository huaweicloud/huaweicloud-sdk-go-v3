package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWarmPoolInstancesResponse Response Object
type ListWarmPoolInstancesResponse struct {
	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 暖池实例列表
	WarmPoolInstances *[]WarmPoolInstance `json:"warm_pool_instances,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o ListWarmPoolInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWarmPoolInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListWarmPoolInstancesResponse", string(data)}, " ")
}
