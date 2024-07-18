package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpGroupsResponse Response Object
type ListIpGroupsResponse struct {

	// 参数解释：IP地址组列表返回对象。
	Ipgroups *[]IpGroup `json:"ipgroups,omitempty"`

	// 参数解释：请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListIpGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListIpGroupsResponse", string(data)}, " ")
}
