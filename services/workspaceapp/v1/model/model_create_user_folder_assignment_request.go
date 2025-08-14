package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserFolderAssignmentRequest Request Object
type CreateUserFolderAssignmentRequest struct {

	// WKS存储ID。
	StorageId string `json:"storage_id"`

	Body *CreateUserFolderReq `json:"body,omitempty"`
}

func (o CreateUserFolderAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserFolderAssignmentRequest struct{}"
	}

	return strings.Join([]string{"CreateUserFolderAssignmentRequest", string(data)}, " ")
}
