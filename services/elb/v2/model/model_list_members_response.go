package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMembersResponse struct {
	// 后端云服务器对象的列表

	Members        *[]MemberResp `json:"members,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMembersResponse struct{}"
	}

	return strings.Join([]string{"ListMembersResponse", string(data)}, " ")
}
