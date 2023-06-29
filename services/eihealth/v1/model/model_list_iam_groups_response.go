package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIamGroupsResponse Response Object
type ListIamGroupsResponse struct {

	// 用户组列表
	Groups *[]IamGroupDto `json:"groups,omitempty"`

	// 总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIamGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIamGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListIamGroupsResponse", string(data)}, " ")
}
