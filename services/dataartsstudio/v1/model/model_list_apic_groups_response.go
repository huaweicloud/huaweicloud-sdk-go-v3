package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApicGroupsResponse Response Object
type ListApicGroupsResponse struct {

	// 网关分组列表
	GatewayGroups  *[]ApigGroupDto `json:"gateway_groups,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListApicGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApicGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListApicGroupsResponse", string(data)}, " ")
}
