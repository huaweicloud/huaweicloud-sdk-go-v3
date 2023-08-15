package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronListSubnetsResponse Response Object
type NeutronListSubnetsResponse struct {

	// subnet对象列表
	Subnets *[]NeutronSubnet `json:"subnets,omitempty"`

	// 分页信息
	SubnetsLinks   *[]NeutronPageLink `json:"subnets_links,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o NeutronListSubnetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListSubnetsResponse struct{}"
	}

	return strings.Join([]string{"NeutronListSubnetsResponse", string(data)}, " ")
}
