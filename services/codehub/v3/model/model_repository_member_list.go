package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryMemberList struct {

	// 仓库成员总数
	Total *int32 `json:"total,omitempty"`

	// 仓库成员列表
	Users *[]RepositoryMember `json:"users,omitempty"`
}

func (o RepositoryMemberList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryMemberList struct{}"
	}

	return strings.Join([]string{"RepositoryMemberList", string(data)}, " ")
}
