package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCertTagRequest struct {

	// 所需要创建标签的证书ID。
	CertificateId string `json:"certificate_id"`

	Body *ResourceTagRequestBody `json:"body,omitempty"`
}

func (o CreateCertTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertTagRequest struct{}"
	}

	return strings.Join([]string{"CreateCertTagRequest", string(data)}, " ")
}
