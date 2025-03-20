package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupsV5Response Response Object
type ListGroupsV5Response struct {

	// 用户组列表。
	Groups *[]Group `json:"groups,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListGroupsV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupsV5Response struct{}"
	}

	return strings.Join([]string{"ListGroupsV5Response", string(data)}, " ")
}
