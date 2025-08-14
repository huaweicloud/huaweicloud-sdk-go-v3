package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCloudStorageAttachmentResponse Response Object
type DeleteCloudStorageAttachmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCloudStorageAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudStorageAttachmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteCloudStorageAttachmentResponse", string(data)}, " ")
}
