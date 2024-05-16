package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateHostsResponse Response Object
type BatchCreateHostsResponse struct {

	// 创建的主机列表
	Items *[]HostItemWithId `json:"items,omitempty"`

	// 主机总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCreateHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateHostsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateHostsResponse", string(data)}, " ")
}
