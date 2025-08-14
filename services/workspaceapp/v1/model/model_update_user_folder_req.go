package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserFolderReq 修改个人文件夹。
type UpdateUserFolderReq struct {
	UpdateUserAssignmentInfo *UpdateUserAssignmentInfo `json:"update_user_assignment_info,omitempty"`
}

func (o UpdateUserFolderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserFolderReq struct{}"
	}

	return strings.Join([]string{"UpdateUserFolderReq", string(data)}, " ")
}
