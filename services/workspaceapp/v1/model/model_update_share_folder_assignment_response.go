package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateShareFolderAssignmentResponse Response Object
type UpdateShareFolderAssignmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateShareFolderAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateShareFolderAssignmentResponse struct{}"
	}

	return strings.Join([]string{"UpdateShareFolderAssignmentResponse", string(data)}, " ")
}
