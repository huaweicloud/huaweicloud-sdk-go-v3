package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserFolderAssignmentResponse Response Object
type UpdateUserFolderAssignmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateUserFolderAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserFolderAssignmentResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserFolderAssignmentResponse", string(data)}, " ")
}
