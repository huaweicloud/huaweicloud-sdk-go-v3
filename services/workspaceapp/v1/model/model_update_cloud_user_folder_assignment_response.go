package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCloudUserFolderAssignmentResponse Response Object
type UpdateCloudUserFolderAssignmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCloudUserFolderAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudUserFolderAssignmentResponse struct{}"
	}

	return strings.Join([]string{"UpdateCloudUserFolderAssignmentResponse", string(data)}, " ")
}
