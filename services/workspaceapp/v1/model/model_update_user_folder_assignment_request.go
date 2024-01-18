package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserFolderAssignmentRequest Request Object
type UpdateUserFolderAssignmentRequest struct {

	// WKS存储ID。
	StorageId string `json:"storage_id"`

	Body *AssignUserFolderReq `json:"body,omitempty"`
}

func (o UpdateUserFolderAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserFolderAssignmentRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserFolderAssignmentRequest", string(data)}, " ")
}
