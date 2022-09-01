package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListGroupsResponse struct {

	// 存储组列表
	Groups *[]GetGroup `json:"groups,omitempty" xml:"groups"`

	// 返回的 data-store-group 数量
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupsResponse", string(data)}, " ")
}
