package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShareFolderRequest Request Object
type CreateShareFolderRequest struct {

	// WKS存储ID
	StorageId string `json:"storage_id"`

	Body *CreateShareFolderReq `json:"body,omitempty"`
}

func (o CreateShareFolderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareFolderRequest struct{}"
	}

	return strings.Join([]string{"CreateShareFolderRequest", string(data)}, " ")
}
