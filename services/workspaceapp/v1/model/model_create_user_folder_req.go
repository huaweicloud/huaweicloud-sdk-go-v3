package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserFolderReq 创建个人文件夹。
type CreateUserFolderReq struct {

	// 存储分配目标。
	Items []CreateUserAssignmentInfo `json:"items"`
}

func (o CreateUserFolderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserFolderReq struct{}"
	}

	return strings.Join([]string{"CreateUserFolderReq", string(data)}, " ")
}
