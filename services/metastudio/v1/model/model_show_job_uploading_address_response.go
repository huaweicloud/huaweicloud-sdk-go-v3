package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobUploadingAddressResponse Response Object
type ShowJobUploadingAddressResponse struct {
	SegmentUrl *ShowJobUploadingAddressRspSegmentUrl `json:"segment_url,omitempty"`

	PackageUrl *ShowJobUploadingAddressRspPackageUrl `json:"package_url,omitempty"`

	// 授权书的上传地址。
	AuthorizationLetterUploadingUrl *string `json:"authorization_letter_uploading_url,omitempty"`
	HttpStatusCode                  int     `json:"-"`
}

func (o ShowJobUploadingAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobUploadingAddressResponse struct{}"
	}

	return strings.Join([]string{"ShowJobUploadingAddressResponse", string(data)}, " ")
}
