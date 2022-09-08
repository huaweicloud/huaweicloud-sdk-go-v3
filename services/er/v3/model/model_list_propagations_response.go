package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPropagationsResponse struct {

	// 路由传播列表
	Propagations *[]Propagation `json:"propagations,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPropagationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPropagationsResponse struct{}"
	}

	return strings.Join([]string{"ListPropagationsResponse", string(data)}, " ")
}
