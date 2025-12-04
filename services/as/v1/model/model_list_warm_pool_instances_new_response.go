package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWarmPoolInstancesNewResponse Response Object
type ListWarmPoolInstancesNewResponse struct {
	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 暖池实例列表
	WarmPoolInstances *[]WarmPoolInstance `json:"warm_pool_instances,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o ListWarmPoolInstancesNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWarmPoolInstancesNewResponse struct{}"
	}

	return strings.Join([]string{"ListWarmPoolInstancesNewResponse", string(data)}, " ")
}
