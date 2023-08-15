package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateShareFolderAssignmentRequest Request Object
type UpdateShareFolderAssignmentRequest struct {

	// WKS存储ID
	StorageId string `json:"storage_id"`

	Body *AssignShareFolderReq `json:"body,omitempty"`
}

func (o UpdateShareFolderAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateShareFolderAssignmentRequest struct{}"
	}

	return strings.Join([]string{"UpdateShareFolderAssignmentRequest", string(data)}, " ")
}
