package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserFolderAssignmentResponse Response Object
type CreateUserFolderAssignmentResponse struct {

	// 创建文件夹返回信息。
	Items          *[]CloudStorageAssignmentCreateInfo `json:"items,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o CreateUserFolderAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserFolderAssignmentResponse struct{}"
	}

	return strings.Join([]string{"CreateUserFolderAssignmentResponse", string(data)}, " ")
}
