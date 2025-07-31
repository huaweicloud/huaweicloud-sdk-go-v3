package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddMembersRequestBody 批量添加镜像成员body
type BatchAddMembersRequestBody struct {

	// 镜像ID列表
	Images []string `json:"images"`

	// 项目ID列表
	Projects []string `json:"projects"`

	// 账号ID列表
	Domains *[]string `json:"domains,omitempty"`

	// 组织URN列表
	Organizations *[]string `json:"organizations,omitempty"`
}

func (o BatchAddMembersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddMembersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddMembersRequestBody", string(data)}, " ")
}
