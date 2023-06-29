package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceGroupsResponse Response Object
type ListResourceGroupsResponse struct {

	// 资源分组总数
	Count *int32 `json:"count,omitempty"`

	// 资源分组列表
	ResourceGroups *[]OneResourceGroupResp `json:"resource_groups,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListResourceGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceGroupsResponse", string(data)}, " ")
}
