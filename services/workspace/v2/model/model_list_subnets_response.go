package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubnetsResponse Response Object
type ListSubnetsResponse struct {

	// 子网列表。
	Subnets *[]SubnetInfo `json:"subnets,omitempty"`

	// 子网数量。
	SubnetSize     *int32 `json:"subnet_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSubnetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetsResponse struct{}"
	}

	return strings.Join([]string{"ListSubnetsResponse", string(data)}, " ")
}
