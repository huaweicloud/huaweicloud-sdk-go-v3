package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStaticRoutesResponse Response Object
type ListStaticRoutesResponse struct {

	// 路由列表
	Routes *[]Route `json:"routes,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListStaticRoutesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStaticRoutesResponse struct{}"
	}

	return strings.Join([]string{"ListStaticRoutesResponse", string(data)}, " ")
}
