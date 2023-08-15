package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssignUserFolderReq 为用户直接自动创建和分配WSK存储目录
type AssignUserFolderReq struct {

	// 存储分配目标
	Items []UserAssignment `json:"items"`
}

func (o AssignUserFolderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssignUserFolderReq struct{}"
	}

	return strings.Join([]string{"AssignUserFolderReq", string(data)}, " ")
}
