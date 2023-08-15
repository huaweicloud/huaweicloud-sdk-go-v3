package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostsResponse Response Object
type ListHostsResponse struct {

	// 主机参数。
	Hosts *[]HostModel `json:"hosts,omitempty"`

	// 主机列表总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsResponse struct{}"
	}

	return strings.Join([]string{"ListHostsResponse", string(data)}, " ")
}
