package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostsDetailResponse Response Object
type ListHostsDetailResponse struct {

	// 云办公主机列表。
	DedicatedHosts *[]ListHostsRspDedicatedHosts `json:"dedicated_hosts,omitempty"`

	// 总共条数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListHostsDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsDetailResponse struct{}"
	}

	return strings.Join([]string{"ListHostsDetailResponse", string(data)}, " ")
}
