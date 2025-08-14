package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCloudUserFolderAssignmentRequest Request Object
type UpdateCloudUserFolderAssignmentRequest struct {

	// WKS存储ID。
	StorageId string `json:"storage_id"`

	// 文件夹id。
	CloudAssignmentId string `json:"cloud_assignment_id"`

	Body *UpdateUserFolderReq `json:"body,omitempty"`
}

func (o UpdateCloudUserFolderAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudUserFolderAssignmentRequest struct{}"
	}

	return strings.Join([]string{"UpdateCloudUserFolderAssignmentRequest", string(data)}, " ")
}
