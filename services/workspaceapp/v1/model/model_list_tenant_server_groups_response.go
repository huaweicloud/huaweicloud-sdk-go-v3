package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantServerGroupsResponse Response Object
type ListTenantServerGroupsResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 服务器组列表。
	Items          *[]ServerGroupDto `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListTenantServerGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantServerGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListTenantServerGroupsResponse", string(data)}, " ")
}
