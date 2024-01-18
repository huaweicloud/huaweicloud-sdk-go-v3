package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserStorageAttachmentRequest Request Object
type DeleteUserStorageAttachmentRequest struct {

	// WKS存储ID。
	StorageId string `json:"storage_id"`

	Body *DeleteUserStorageAttachmentReq `json:"body,omitempty"`
}

func (o DeleteUserStorageAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserStorageAttachmentRequest struct{}"
	}

	return strings.Join([]string{"DeleteUserStorageAttachmentRequest", string(data)}, " ")
}
