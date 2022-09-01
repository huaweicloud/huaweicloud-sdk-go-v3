package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRepoMemberResult struct {

	// 用户id
	Id *string `json:"id,omitempty" xml:"id"`

	// 创建仓库成员信息
	Message *string `json:"message,omitempty" xml:"message"`

	// 用户名
	Name *string `json:"name,omitempty" xml:"name"`

	// 创建仓库成员状态
	Status *string `json:"status,omitempty" xml:"status"`
}

func (o CreateRepoMemberResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepoMemberResult struct{}"
	}

	return strings.Join([]string{"CreateRepoMemberResult", string(data)}, " ")
}
