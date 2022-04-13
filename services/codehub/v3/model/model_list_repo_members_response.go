package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRepoMembersResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *RepositoryMemberList `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepoMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepoMembersResponse struct{}"
	}

	return strings.Join([]string{"ListRepoMembersResponse", string(data)}, " ")
}
