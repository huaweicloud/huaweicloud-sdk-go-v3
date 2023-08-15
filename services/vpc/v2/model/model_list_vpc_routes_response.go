package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpcRoutesResponse Response Object
type ListVpcRoutesResponse struct {

	// route对象列表
	Routes *[]VpcRoute `json:"routes,omitempty"`

	// 分页信息
	RoutesLinks    *[]NeutronPageLink `json:"routes_links,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListVpcRoutesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcRoutesResponse struct{}"
	}

	return strings.Join([]string{"ListVpcRoutesResponse", string(data)}, " ")
}
