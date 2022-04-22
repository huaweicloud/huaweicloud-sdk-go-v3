package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSubnetsResponse struct {

	// 子网数组。
	Subnets *[]Subnet `json:"subnets,omitempty"`

	// 子网数目。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSubnetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetsResponse struct{}"
	}

	return strings.Join([]string{"ListSubnetsResponse", string(data)}, " ")
}
