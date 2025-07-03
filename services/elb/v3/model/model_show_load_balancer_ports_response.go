package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoadBalancerPortsResponse Response Object
type ShowLoadBalancerPortsResponse struct {

	// 当前ELB占用的ports列表。
	Ports *[]LocalPort `json:"ports,omitempty"`

	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLoadBalancerPortsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadBalancerPortsResponse struct{}"
	}

	return strings.Join([]string{"ShowLoadBalancerPortsResponse", string(data)}, " ")
}
