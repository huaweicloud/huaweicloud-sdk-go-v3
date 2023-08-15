package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserStorageAttachmentResponse Response Object
type DeleteUserStorageAttachmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteUserStorageAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserStorageAttachmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteUserStorageAttachmentResponse", string(data)}, " ")
}
