package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadFilePublisherRequest Request Object
type UploadFilePublisherRequest struct {

	// 发布商凭证,x-publisher-token和X-Auth-Token必传一个
	XPublisherToken *string `json:"x-publisher-token,omitempty"`

	Body *UploadFilePublisherRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadFilePublisherRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFilePublisherRequest struct{}"
	}

	return strings.Join([]string{"UploadFilePublisherRequest", string(data)}, " ")
}
