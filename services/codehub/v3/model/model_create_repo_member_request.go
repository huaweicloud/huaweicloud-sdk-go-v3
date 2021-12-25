package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRepoMemberRequest struct {
	// 添加用户的信息列表

	Users *[]RepoMemberInfo `json:"users,omitempty"`
}

func (o CreateRepoMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepoMemberRequest struct{}"
	}

	return strings.Join([]string{"CreateRepoMemberRequest", string(data)}, " ")
}
