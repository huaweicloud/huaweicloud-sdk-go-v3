package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIamGroupUsersResponse Response Object
type ListIamGroupUsersResponse struct {

	// 用户组列表
	Groups *[]IamGroupDto `json:"groups,omitempty"`

	// 总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIamGroupUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIamGroupUsersResponse struct{}"
	}

	return strings.Join([]string{"ListIamGroupUsersResponse", string(data)}, " ")
}
