package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReservedInstanceConfigsResponse Response Object
type ListReservedInstanceConfigsResponse struct {

	// 函数预留实例列表
	ReservedInstances *[]ReservedInstanceConfigs `json:"reserved_instances,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 函数个数
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListReservedInstanceConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReservedInstanceConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListReservedInstanceConfigsResponse", string(data)}, " ")
}
