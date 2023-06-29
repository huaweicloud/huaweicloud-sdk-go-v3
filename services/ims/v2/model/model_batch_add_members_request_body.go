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
}

func (o BatchAddMembersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddMembersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddMembersRequestBody", string(data)}, " ")
}
