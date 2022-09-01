package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDedicatedHostsResponse struct {

	// 满足查询条件的专属主机。
	DedicatedHosts *[]RespDedicatedHost `json:"dedicated_hosts,omitempty" xml:"dedicated_hosts"`

	// 满足查询条件的专属主机数量。
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDedicatedHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedHostsResponse struct{}"
	}

	return strings.Join([]string{"ListDedicatedHostsResponse", string(data)}, " ")
}
