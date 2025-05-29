package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificateTagRequest Request Object
type CreateCertificateTagRequest struct {

	// 证书id。
	ResourceId string `json:"resource_id"`

	Body *CreateCertificateTagRequestBody `json:"body,omitempty"`
}

func (o CreateCertificateTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateTagRequest struct{}"
	}

	return strings.Join([]string{"CreateCertificateTagRequest", string(data)}, " ")
}
