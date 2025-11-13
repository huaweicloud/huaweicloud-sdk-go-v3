package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserReviewAttachmentUploadingAddressResponse Response Object
type ShowUserReviewAttachmentUploadingAddressResponse struct {

	// 地址列表
	Addresses      *[]AttachmentUploadingAddress `json:"addresses,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowUserReviewAttachmentUploadingAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserReviewAttachmentUploadingAddressResponse struct{}"
	}

	return strings.Join([]string{"ShowUserReviewAttachmentUploadingAddressResponse", string(data)}, " ")
}
