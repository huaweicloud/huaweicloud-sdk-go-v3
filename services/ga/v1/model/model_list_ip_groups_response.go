package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpGroupsResponse Response Object
type ListIpGroupsResponse struct {

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	IpGroups *[]IpGroupDetail `json:"ip_groups,omitempty"`

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
