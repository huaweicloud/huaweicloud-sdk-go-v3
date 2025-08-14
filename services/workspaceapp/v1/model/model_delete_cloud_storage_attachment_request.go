package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCloudStorageAttachmentRequest Request Object
type DeleteCloudStorageAttachmentRequest struct {

	// WKS存储ID。
	StorageId string `json:"storage_id"`

	Body *DeleteCloudStorageAttachmentReq `json:"body,omitempty"`
}

func (o DeleteCloudStorageAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudStorageAttachmentRequest struct{}"
	}

	return strings.Join([]string{"DeleteCloudStorageAttachmentRequest", string(data)}, " ")
}
