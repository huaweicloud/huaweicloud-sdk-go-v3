package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDirectConnectsResponse Response Object
type ListDirectConnectsResponse struct {

	// 操作请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 物理专线对象列表
	DirectConnects *[]DirectConnect `json:"direct_connects,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDirectConnectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectConnectsResponse struct{}"
	}

	return strings.Join([]string{"ListDirectConnectsResponse", string(data)}, " ")
}
