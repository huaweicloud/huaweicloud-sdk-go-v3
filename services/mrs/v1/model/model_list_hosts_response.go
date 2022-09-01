package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHostsResponse struct {

	// 主机参数。
	Hosts *[]HostModel `json:"hosts,omitempty" xml:"hosts"`

	// 主机列表总数。
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsResponse struct{}"
	}

	return strings.Join([]string{"ListHostsResponse", string(data)}, " ")
}
