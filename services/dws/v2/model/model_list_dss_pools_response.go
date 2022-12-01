package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDssPoolsResponse struct {

	// 专属分布式存储池详情列表。
	Pools *[]DssPool `json:"pools,omitempty"`

	// 专属分布式存储池个数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDssPoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDssPoolsResponse struct{}"
	}

	return strings.Join([]string{"ListDssPoolsResponse", string(data)}, " ")
}
