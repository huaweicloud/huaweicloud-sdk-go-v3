package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowShardingBalancerResponse struct {

	// 集群均衡是否打开。
	IsOpen *bool `json:"is_open,omitempty" xml:"is_open"`

	ActiveWindow   *BalancerActiveWindow `json:"active_window,omitempty" xml:"active_window"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowShardingBalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShardingBalancerResponse struct{}"
	}

	return strings.Join([]string{"ShowShardingBalancerResponse", string(data)}, " ")
}
