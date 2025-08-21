package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRepositoryMembersResponse Response Object
type AddRepositoryMembersResponse struct {

	// 批量创建成员响应
	Body           *[]BatchCreateRepositoryMemberDto `json:"body,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o AddRepositoryMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRepositoryMembersResponse struct{}"
	}

	return strings.Join([]string{"AddRepositoryMembersResponse", string(data)}, " ")
}
